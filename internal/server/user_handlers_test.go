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

	client := pb.NewUserServiceClient(conn)

	t.Run("create_user", func(t *testing.T) {
		_, err := client.CreateUser(ctx, &pb.NewUserReq{
			Login:    "I0HuKc",
			Email:    "kiwi@gmail.com",
			Password: "somepasswordhere5",
		})

		assert.NoError(t, err)
	})
}
