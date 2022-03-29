package mocksqlstore

import (
	"database/sql"
	"time"

	"github.com/wrs-news/bfb-user-microservice/internal/models"
	pb "github.com/wrs-news/golang-proto/pkg/proto/user"
)

type UserRepository struct {
	users map[int32]*models.User
}

func (r *UserRepository) Create(u *models.User) error {
	u.CreatedAt = time.Now().UTC().Format("2006-01-02T15:04:05.00000000")
	u.UpdatedAt = time.Now().UTC().Format("2006-01-02T15:04:05.00000000")

	r.users[u.Id] = u
	return nil
}

func (r *UserRepository) GetById(u *models.User) error {
	for k, v := range r.users {
		if v.Id == k {
			u = v
			return nil
		}
	}

	return sql.ErrNoRows
}

func (r *UserRepository) GetByUuid(u *models.User) error {
	for _, v := range r.users {
		if v.Uuid == u.Uuid {
			u = v
			return nil
		}
	}

	return sql.ErrNoRows
}

func (r *UserRepository) GetByLogin(u *models.User) error {
	for _, v := range r.users {
		if v.Login == u.Login {
			u = v
			return nil
		}
	}

	return sql.ErrNoRows
}

func (r *UserRepository) Update(u *models.User) error {
	for _, v := range r.users {
		if v.Uuid == u.Uuid {
			u = v
			return nil
		}
	}

	return sql.ErrNoRows
}

func (r *UserRepository) Delete(u *models.User) error {
	for _, v := range r.users {
		if v.Uuid == u.Uuid {
			u = v
			delete(r.users, v.Id)
			return nil
		}
	}
	return sql.ErrNoRows
}

func (r *UserRepository) Count(sReq *pb.SelectionReq) (int32, error) {
	return int32(len(r.users)), nil
}

func (r *UserRepository) Selection(sReq *pb.SelectionReq) ([]*pb.User, error) {
	return nil, nil
}
