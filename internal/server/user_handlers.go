package server

import (
	"context"

	"github.com/wrs-news/bfb-user-microservice/internal/models"
	pb "github.com/wrs-news/golang-proto/pkg/proto/user"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) CreateUser(ctx context.Context, nu *pb.NewUser) (*pb.User, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	u := models.User{
		Login: nu.Login,
		Email: nu.Email,
		Hash:  string(b),
	}

	if err := s.sqlstore.User().Create(&u); err != nil {
		return nil, err
	}

	return &pb.User{
		Id:        u.Id,
		Uuid:      u.Uuid,
		Login:     u.Login,
		Email:     u.Email,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}
