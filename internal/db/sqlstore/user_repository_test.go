package sqlstore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wrs-news/bfb-user-microservice/internal/db"
	"github.com/wrs-news/bfb-user-microservice/internal/db/sqlstore"
	"github.com/wrs-news/bfb-user-microservice/internal/models"
	pb "github.com/wrs-news/golang-proto/pkg/proto/user"
)

func Test_DB_UserRepository(t *testing.T) {
	db, teardown := db.TestDB(t, &TestConfig.Services.DB)
	defer teardown("users")

	s := sqlstore.Create(db)

	u := models.User{
		Login: "I0HuKc",
		Email: "kiwi@gmail.com",
		Hash:  "somehash",
	}

	// Создание пользователя
	t.Run("create_user", func(t *testing.T) {
		assert.NoError(t, s.User().Create(&u))
	})

	// Обновление роли пользователя
	t.Run("update_user", func(t *testing.T) {
		u.Role = 1
		assert.NoError(t, s.User().Update(&u))
	})

	// Получение пользователя по Id
	t.Run("get_user_by_id", func(t *testing.T) {
		assert.NoError(t, s.User().GetById(&u))
	})

	// Получение пользователя по Uuid
	t.Run("get_user_by_uuid", func(t *testing.T) {
		assert.NoError(t, s.User().GetByUuid(&u))
	})

	// Получение пользователя по Login
	t.Run("get_user_by_login", func(t *testing.T) {
		assert.NoError(t, s.User().GetByLogin(&u))
	})

	// Проверка обновления роли
	assert.Equal(t, int32(1), u.Role)

	sReq := pb.SelectionReq{
		Limit:  15,
		Offset: 0,
	}

	// Получение кол-ва записей
	t.Run("count_records", func(t *testing.T) {
		c, err := s.User().Count(&sReq)
		assert.NoError(t, err)
		assert.Equal(t, int32(1), c)
	})

	// Получение массива записей
	t.Run("get_user_records", func(t *testing.T) {
		arr, err := s.User().Selection(&sReq)
		assert.NoError(t, err)
		assert.Len(t, arr, int(1))
		assert.IsType(t, []*pb.User{}, arr)
	})

	// Удаление пользователя
	t.Run("delete_user", func(t *testing.T) {
		assert.NoError(t, s.User().Delete(&u))
		assert.Error(t, s.User().GetById(&u))
	})

}
