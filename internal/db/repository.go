package db

import "github.com/wrs-news/bfb-user-microservice/internal/models"

type UserRepositoryI interface {
	Create(u *models.User) error
}
