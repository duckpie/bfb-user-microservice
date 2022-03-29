package server_test

import (
	"context"
	"log"
	"net"
	"os"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/wrs-news/bfb-user-microservice/internal/config"
	"github.com/wrs-news/bfb-user-microservice/internal/db/mocksqlstore"
	"github.com/wrs-news/bfb-user-microservice/internal/server"
	pb "github.com/wrs-news/golang-proto/pkg/proto/user"
	"google.golang.org/grpc/test/bufconn"
)

var (
	testConfig = config.NewConfig()
	bufSize    = 1024 * 1024
	lis        *bufconn.Listener
)

func TestMain(m *testing.M) {
	if _, err := toml.DecodeFile("../../config/config.local.toml", testConfig); err != nil {
		log.Fatalf(err.Error())
	}

	// Подменяю ссылку на тестовую БД
	testConfig.Services.DB.DbUrl = testConfig.Services.DB.DbUrlTest

	lis = bufconn.Listen(bufSize)
	srv := server.InitServer(mocksqlstore.Create())

	pb.RegisterUserServiceServer(srv.GetServer(), srv)
	go func() {
		if err := srv.GetServer().Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	os.Exit(m.Run())
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}
