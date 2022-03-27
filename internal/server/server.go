package server

import (
	"github.com/wrs-news/bfb-user-microservice/internal/db"
	pb "github.com/wrs-news/golang-proto/pkg/proto/user"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedUserServiceServer

	server   *grpc.Server
	sqlstore db.SQLStoreI
}

type ServerI interface {
	GetServer() *grpc.Server
}

func (s *Server) GetServer() *grpc.Server {
	return s.server
}

func InitServer(s db.SQLStoreI) *Server {
	return &Server{
		server:   grpc.NewServer(),
		sqlstore: s,
	}
}
