package sqlstore

import (
	"database/sql"

	"github.com/wrs-news/bfb-user-microservice/internal/models"
	pb "github.com/wrs-news/golang-proto/pkg/proto/user"
)

type UserRepository struct {
	store *sql.DB
}

func (r *UserRepository) Create(u *models.User) error {
	return r.store.QueryRow(
		`
		INSERT INTO users(login, email, hash)
		SELECT $1, $2, $3
		RETURNING id, uuid, login, email, hash, role, created_at, updated_at
		`,
		u.Login,
		u.Email,
		u.Hash,
	).Scan(
		&u.Id,
		&u.Uuid,
		&u.Login,
		&u.Email,
		&u.Hash,
		&u.Role,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
}

func (r *UserRepository) GetById(u *models.User) error {
	return r.store.QueryRow(
		`
		SELECT uuid, login, email, hash, role, created_at, updated_at
		FROM users
		WHERE id=$1
		`,
		u.Id,
	).Scan(
		&u.Uuid,
		&u.Login,
		&u.Email,
		&u.Hash,
		&u.Role,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
}

func (r *UserRepository) GetByUuid(u *models.User) error {
	return r.store.QueryRow(
		`
		SELECT id, login, email, hash, role, created_at, updated_at
		FROM users
		WHERE uuid=$1
		`,
		u.Uuid,
	).Scan(
		&u.Id,
		&u.Login,
		&u.Email,
		&u.Hash,
		&u.Role,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
}

func (r *UserRepository) GetByLogin(u *models.User) error {
	return r.store.QueryRow(
		`
		SELECT id, uuid, email, hash, role, created_at, updated_at
		FROM users
		WHERE login=$1
		`,
		u.Login,
	).Scan(
		&u.Id,
		&u.Uuid,
		&u.Email,
		&u.Hash,
		&u.Role,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
}

func (r *UserRepository) Delete(u *models.User) error {
	return r.store.QueryRow(
		`
		DELETE FROM users
		WHERE uuid=$1
		RETURNING id, login, email, hash, role, created_at, updated_at
		`,
		u.Uuid,
	).Scan(
		&u.Id,
		&u.Login,
		&u.Email,
		&u.Hash,
		&u.Role,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
}

func (r *UserRepository) Update(u *models.User) error {
	return r.store.QueryRow(
		`
		UPDATE users 
		SET login=$1, email=$2, role=$3
		WHERE uuid=$4
		RETURNING id, login, email, hash, role, created_at, updated_at
		`,
		u.Login,
		u.Email,
		u.Role,
		u.Uuid,
	).Scan(
		&u.Id,
		&u.Login,
		&u.Email,
		&u.Hash,
		&u.Role,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
}

func (r *UserRepository) Count(sReq *pb.SelectionReq) (int32, error) {
	var c int32
	if err := r.store.QueryRow(
		`
		SELECT count(*)
		FROM users		
		`,
	).Scan(
		&c,
	); err != nil {
		return 0, err
	}

	return c, nil
}

func (r *UserRepository) Selection(sReq *pb.SelectionReq) ([]*pb.User, error) {
	arr := []*pb.User{}

	rows, err := r.store.Query(
		`
		SELECT  id, uuid, login, email, role, created_at, updated_at
		FROM users
		ORDER BY id DESC
		OFFSET $1
		LIMIT $2
		`,
		sReq.Offset,
		sReq.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		u := pb.User{}
		if err := rows.Scan(
			&u.Id,
			&u.Uuid,
			&u.Login,
			&u.Email,
			&u.Role,
			&u.CreatedAt,
			&u.UpdatedAt,
		); err != nil {
			continue
		}

		arr = append(arr, &u)
	}

	return arr, nil

}
