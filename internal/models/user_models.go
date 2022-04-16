package models

import (
	pb "github.com/wrs-news/golang-proto/pkg/proto/user"
)

type User struct {
	Id        int32  `json:"id"`
	Uuid      string `json:"uuid"`
	Login     string `json:"login"`
	Email     string `json:"email"`
	Hash      string `json:"hash"`
	Role      int32  `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (u *User) ToProtoUser() *pb.User {
	return &pb.User{
		Id:        u.Id,
		Uuid:      u.Uuid,
		Login:     u.Login,
		Email:     u.Email,
		Hash:      u.Hash,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
