package db

import (
	"database/sql"

	_ "github.com/lib/pq" // database driver
	"github.com/wrs-news/bfb-user-microservice/internal/config"
)

func InitPostgres(config *config.DatabaseConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", config.DbUrl)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
