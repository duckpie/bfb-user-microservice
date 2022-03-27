package main

// import (
// 	"context"
// 	"flag"
// 	"fmt"
// 	"log"
// 	"net"

// 	pb "github.com/wrs-news/golang-proto/pkg/proto/user"
// 	"google.golang.org/grpc"
// )

// var (
// 	port = flag.Int("port", 50051, "The server port")
// )

// type server struct {
// 	pb.UnimplementedUserServiceServer
// }

// func (s *server) CreateUser(ctx context.Context, ur *pb.NewUser) (*pb.User, error) {
// 	fmt.Printf("111test")
// 	return &pb.User{
// 		Id:        1,
// 		Login:     "login",
// 		Email:     "email",
// 		Role:      1,
// 		CreatedAt: "egwg",
// 		UpdatedAt: "dfdf",
// 	}, nil
// }

// func main() {
// 	flag.Parse()
// 	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}
// 	s := grpc.NewServer()
// 	pb.RegisterUserServiceServer(s, &server{})
// 	log.Printf("server listening at %v", lis.Addr())
// 	if err := s.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }

import (
	"github.com/spf13/cobra"

	"github.com/wrs-news/bfb-user-microservice/cmd"
)

func main() {
	cobra.CheckErr(cmd.NewRootCmd().Execute())
}
