package server

import (
	"context"

	"github.com/wrs-news/bfb-user-microservice/internal/db"
	pb "github.com/wrs-news/golang-proto/pkg/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/encoding/gzip"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (s *Server) HeartbeatCheck(ctx context.Context, e *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func CreateServer(s db.SQLStoreI) *Server {
	encoding.RegisterCompressor(encoding.GetCompressor(gzip.Name))

	return &Server{
		server:   grpc.NewServer(),
		sqlstore: s,
	}
}
