package mocksqlstore

import (
	"github.com/wrs-news/bfb-user-microservice/internal/db"
	"github.com/wrs-news/bfb-user-microservice/internal/models"
)

type Store struct {
	userRepository *UserRepository
}

func Create() db.SQLStoreI {
	return &Store{}
}

func (s *Store) User() db.UserRepositoryI {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		users: make(map[int32]*models.User),
	}

	return s.userRepository
}
