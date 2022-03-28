package db

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"

	"github.com/wrs-news/bfb-user-microservice/internal/config"
)

func TestDB(t *testing.T, config *config.DatabaseConfig) (*sql.DB, func(...string)) {
	t.Helper()

	// Инициализирую БД
	db, err := sql.Open("postgres", config.DbUrlTest)
	if err != nil {
		t.Fatal(err)
	}

	// Поверка подключения к БД
	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			// Очищаю таблицы
			str := fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))
			db.Exec(str)
		}

		db.Close()
	}
}
