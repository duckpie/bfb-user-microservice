package server

import (
	"context"
	"database/sql"

	"github.com/wrs-news/bfb-user-microservice/internal/models"
	pb "github.com/wrs-news/golang-proto/pkg/proto/user"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateUser(ctx context.Context, nu *pb.NewUserReq) (*pb.User, error) {
	// Создание хеша пароля
	b, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.MinCost)
	if err != nil {
		es := status.New(codes.Internal, err.Error())
		return nil, es.Err()
	}

	u := models.User{
		Login: nu.Login,
		Email: nu.Email,
		Hash:  string(b),
	}

	// Операция с БД
	if err := s.sqlstore.User().Create(&u); err != nil {
		es := status.New(codes.Internal, err.Error())
		return nil, es.Err()
	}

	return u.ToProtoUser(), nil
}

func (s *Server) GetUserById(ctx context.Context, uReq *pb.UserReqID) (*pb.User, error) {
	u := models.User{Id: uReq.Id}

	if err := s.sqlstore.User().GetById(&u); err != nil {
		switch err {
		case sql.ErrNoRows:
			es := status.New(codes.NotFound, err.Error())
			return nil, es.Err()

		default:
			es := status.New(codes.Internal, err.Error())
			return nil, es.Err()
		}
	}

	return u.ToProtoUser(), nil
}

func (s *Server) GetUserByLogin(ctx context.Context, uReq *pb.UserReqLogin) (*pb.User, error) {
	u := models.User{Login: uReq.Login}

	if err := s.sqlstore.User().GetByLogin(&u); err != nil {
		switch err {
		case sql.ErrNoRows:
			es := status.New(codes.NotFound, err.Error())
			return nil, es.Err()

		default:
			es := status.New(codes.Internal, err.Error())
			return nil, es.Err()
		}
	}

	return u.ToProtoUser(), nil
}

func (s *Server) GetUserByUuid(ctx context.Context, uReq *pb.UserReqUuid) (*pb.User, error) {
	u := models.User{Uuid: uReq.Uuid}

	if err := s.sqlstore.User().GetByUuid(&u); err != nil {
		switch err {
		case sql.ErrNoRows:
			es := status.New(codes.NotFound, err.Error())
			return nil, es.Err()

		default:
			es := status.New(codes.Internal, err.Error())
			return nil, es.Err()
		}
	}

	return u.ToProtoUser(), nil
}

func (s *Server) UpdateUser(ctx context.Context, uUreq *pb.UpdateUserReq) (*pb.User, error) {
	u := models.User{
		Uuid:  uUreq.Uuid,
		Login: uUreq.Login,
		Email: uUreq.Email,
		Role:  uUreq.Role,
	}

	if err := s.sqlstore.User().Update(&u); err != nil {
		switch err {
		case sql.ErrNoRows:
			es := status.New(codes.NotFound, err.Error())
			return nil, es.Err()

		default:
			es := status.New(codes.Internal, err.Error())
			return nil, es.Err()
		}
	}

	return u.ToProtoUser(), nil
}

func (s *Server) DeleteUser(ctx context.Context, uReq *pb.UserReqUuid) (*pb.User, error) {
	u := models.User{Uuid: uReq.Uuid}

	if err := s.sqlstore.User().Delete(&u); err != nil {
		switch err {
		case sql.ErrNoRows:
			es := status.New(codes.NotFound, err.Error())
			return nil, es.Err()

		default:
			es := status.New(codes.Internal, err.Error())
			return nil, es.Err()
		}
	}

	return u.ToProtoUser(), nil
}

func (s *Server) GetAll(ctx context.Context, sReq *pb.SelectionReq) (*pb.Selection, error) {
	cArr := make(chan []*pb.User)
	cCount := make(chan int32)
	errs, ctx := errgroup.WithContext(ctx)

	errs.Go(func() error {
		defer close(cArr)
		arr, err := s.sqlstore.User().Selection(sReq)
		if err != nil {
			es := status.New(codes.Internal, err.Error())
			return es.Err()
		}

		cArr <- arr
		return nil
	})

	errs.Go(func() error {
		defer close(cCount)
		c, err := s.sqlstore.User().Count(sReq)
		if err != nil {
			es := status.New(codes.Internal, err.Error())
			return es.Err()
		}

		cCount <- c
		return nil
	})

	c := <-cCount
	arr := <-cArr

	if errs.Wait() != nil {
		es := status.New(codes.Internal, errs.Wait().Error())
		return nil, es.Err()
	}

	return &pb.Selection{
		Limit:    sReq.Limit,
		Offset:   sReq.Offset,
		Total:    c,
		Data:     arr,
		LastPage: 32,
	}, nil
}
