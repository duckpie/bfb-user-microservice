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
	testCases := []struct {
		name    string
		payload pb.NewUserReq
		isValid bool
	}{
		{
			name: "empty_login",
			payload: pb.NewUserReq{
				Email:    "kiwi@gmail.com",
				Password: "somepasswordhere1",
			},
			isValid: false,
		},
		{
			name: "empty_email",
			payload: pb.NewUserReq{
				Login:    "I0HuKc",
				Password: "somepasswordhere2",
			},
			isValid: false,
		},
		{
			name: "empty_password",
			payload: pb.NewUserReq{
				Login: "I0HuKc",
				Email: "kiwi@gmail.com",
			},
			isValid: true,
		},
		{
			name: "invalid_email",
			payload: pb.NewUserReq{
				Login:    "I0HuKc",
				Email:    "invalid",
				Password: "somepasswordhere4",
			},
			isValid: false,
		},
		{
			name: "valid",
			payload: pb.NewUserReq{
				Login:    "I0HuKc",
				Email:    "kiwi@gmail.com",
				Password: "somepasswordhere5",
			},
			isValid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := client.CreateUser(ctx, &tc.payload)

			if tc.isValid {
				assert.NoError(t, err)
				return
			}

			assert.Error(t, err)
		})
	}

}
