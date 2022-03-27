package sqlstore

import (
	"database/sql"

	"github.com/wrs-news/bfb-user-microservice/internal/models"
)

type UserRepository struct {
	store *sql.DB
}

func (r *UserRepository) Create(u *models.User) error {
	if err := r.store.QueryRow(
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
	); err != nil {
		return err
	}

	return nil
}
