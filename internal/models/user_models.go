package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
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
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func NewUserReq_Validation(nu *pb.NewUserReq) error {
	return validation.ValidateStruct(
		nu,

		validation.Field(&nu.Email,
			is.Email,
			validation.Required,
		),

		validation.Field(&nu.Login,
			validation.Required,
		),

		validation.Field(&nu.Password,
			validation.Required,
		),
	)
}
