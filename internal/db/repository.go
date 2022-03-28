package db

import (
	"github.com/wrs-news/bfb-user-microservice/internal/models"
	pb "github.com/wrs-news/golang-proto/pkg/proto/user"
)

type UserRepositoryI interface {
	Create(u *models.User) error
	GetById(u *models.User) error
	GetByUuid(u *models.User) error
	GetByLogin(u *models.User) error
	Update(u *models.User) error
	Delete(u *models.User) error

	Count(sReq *pb.SelectionReq) (int32, error)
	Selection(sReq *pb.SelectionReq) ([]*pb.User, error)
}
