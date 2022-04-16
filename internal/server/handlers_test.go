package server_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	pb "github.com/wrs-news/golang-proto/pkg/proto/user"
	"google.golang.org/grpc"
)

func Test_Server_UserHandlers_CreateUser(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	var user *pb.User
	client := pb.NewUserServiceClient(conn)

	t.Run("create_user", func(t *testing.T) {
		u, err := client.CreateUser(ctx, &pb.NewUserReq{
			Login:    "I0HuKc",
			Email:    "kiwi@gmail.com",
			Password: "somepasswordhere5",
		})

		assert.NoError(t, err)
		assert.NotNil(t, u)
		user = u
	})

	t.Run("get_user_by_id", func(t *testing.T) {
		u, err := client.GetUserById(ctx, &pb.UserReqID{Id: user.Id})

		assert.NoError(t, err)
		assert.NotNil(t, u)
		assert.Equal(t, user.Id, u.Id)
	})

	t.Run("get_user_by_uuid", func(t *testing.T) {
		u, err := client.GetUserByUuid(ctx, &pb.UserReqUuid{Uuid: user.Uuid})

		assert.NoError(t, err)
		assert.NotNil(t, u)
		assert.Equal(t, user.Uuid, u.Uuid)
	})

	t.Run("get_user_by_login", func(t *testing.T) {
		u, err := client.GetUserByLogin(ctx, &pb.UserReqLogin{Login: user.Login})

		assert.NoError(t, err)
		assert.NotNil(t, u)
		assert.Equal(t, user.Login, u.Login)
	})

	t.Run("user_update", func(t *testing.T) {
		u, err := client.UpdateUser(ctx, &pb.UpdateUserReq{
			Uuid:  user.Uuid,
			Login: user.Login,
			Email: user.Email,
			Role:  1,
		})

		assert.NoError(t, err)
		assert.NotNil(t, u)
		assert.Equal(t, user.Uuid, u.Uuid)
		assert.Equal(t, 1, int(u.Role))
	})

	t.Run("get_users_slice", func(t *testing.T) {
		d, err := client.GetAll(ctx, &pb.SelectionReq{Limit: 15, Offset: 0})

		assert.NoError(t, err)
		assert.NotNil(t, d)
		assert.Len(t, d.Data, 1)

	})

	t.Run("delete_user", func(t *testing.T) {
		u, err := client.DeleteUser(ctx, &pb.UserReqUuid{Uuid: user.Uuid})

		assert.NoError(t, err)
		assert.NotNil(t, u)
		assert.Equal(t, user.Uuid, u.Uuid)
	})
}
