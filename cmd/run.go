package cmd

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/BurntSushi/toml"
	"github.com/oklog/oklog/pkg/group"
	"github.com/spf13/cobra"

	"github.com/wrs-news/bfb-user-microservice/internal/config"
	"github.com/wrs-news/bfb-user-microservice/internal/db"
	"github.com/wrs-news/bfb-user-microservice/internal/db/sqlstore"
	"github.com/wrs-news/bfb-user-microservice/internal/server"
	pb "github.com/wrs-news/golang-proto/pkg/proto/user"
)

var (
	env string
)

func runCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run microservice",
		Long:  `...`,
		Run: func(cmd *cobra.Command, args []string) {
			cfg := config.NewConfig()

			if _, err := toml.DecodeFile(
				fmt.Sprintf("config/config.%s.toml", env), cfg); err != nil {
				log.Printf(err.Error())
				os.Exit(1)
			}

			if err := runner(cfg); err != nil {
				log.Printf(err.Error())
				os.Exit(1)
			}
		},
	}

	flag.StringVar(&env, "env", "local", "Launch environment")

	flag.Parse()
	return cmd
}

func runner(cfg *config.Config) (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("unknown panic")
			}
		}
	}()

	_, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	postgres, err := db.InitPostgres(&cfg.Services.DB)
	if err != nil {
		return err
	}
	defer postgres.Close()

	srv := server.InitServer(sqlstore.Create(postgres))

	var g group.Group
	{

		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Services.Server.Port))
		if err != nil {
			return err
		}
		log.Printf("server listening a t %v", lis.Addr())

		g.Add(func() error {
			pb.RegisterUserServiceServer(srv.GetServer(), srv)
			return srv.GetServer().Serve(lis)
		}, func(error) {
			lis.Close()
		})
	}

	{
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}

	return g.Run()
}
