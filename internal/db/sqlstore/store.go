package sqlstore

import (
	"database/sql"

	"github.com/wrs-news/bfb-user-microservice/internal/db"
)

type Store struct {
	db *sql.DB

	userRepository *UserRepository
}

func Create(db *sql.DB) db.SQLStoreI {
	return &Store{
		db: db,
	}
}

func (s *Store) User() db.UserRepositoryI {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s.db,
	}

	return s.userRepository
}
