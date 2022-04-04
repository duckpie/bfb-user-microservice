package mocksqlstore

import (
	"database/sql"
	"time"

	"github.com/twinj/uuid"
	"github.com/wrs-news/bfb-user-microservice/internal/models"
	pb "github.com/wrs-news/golang-proto/pkg/proto/user"
)

type UserRepository struct {
	users map[int32]*models.User
}

func (r *UserRepository) Create(u *models.User) error {
	u.Id = int32(len(r.users) + 1)
	u.Uuid = uuid.NewV4().String()
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
			v.Login = u.Login
			v.Email = u.Email
			v.Role = u.Role

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
	arr := []*pb.User{}

	for _, u := range r.users {
		arr = append(arr, &pb.User{
			Id:        u.Id,
			Uuid:      u.Uuid,
			Login:     u.Login,
			Email:     u.Email,
			Role:      u.Role,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		})
	}

	var h int32
	if int(sReq.Limit) > len(r.users) {
		h = int32(len(r.users))
	} else {
		h = sReq.Limit
	}

	return arr[sReq.Offset:(sReq.Offset + h)], nil
}

func (r *UserRepository) overwriting(from *models.User, to *models.User) *models.User {
	to.Id = from.Id
	to.Uuid = from.Uuid
	to.Login = from.Login
	to.Email = from.Email
	to.Hash = from.Hash
	to.CreatedAt = from.CreatedAt
	to.UpdatedAt = from.UpdatedAt
	return to
}
