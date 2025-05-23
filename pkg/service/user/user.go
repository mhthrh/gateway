package user

import (
	"context"
	"fmt"
	"github.com/mhthrh/common_pkg/pkg/logger"
	usr "github.com/mhthrh/common_pkg/pkg/model/user"
	userGrpc "github.com/mhthrh/common_pkg/pkg/model/user/grpc/v1"
	"github.com/mhthrh/common_pkg/pkg/xErrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type Service struct {
	l logger.ILogger
}

func New(log logger.ILogger) *Service {
	return &Service{
		l: log,
	}
}
func (s *Service) Create(ctx context.Context, u *usr.User, c *grpc.ClientConn) *xErrors.Error {
	cnn := userGrpc.NewUserServiceClient(c)
	e, stat := cnn.Create(context.Background(), &userGrpc.UserRequest{
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
		UserName:    u.UserName,
		Password:    u.Password,
	})
	st, ok := status.FromError(stat)
	fmt.Println(e, st, ok)
	return nil
}

func (s *Service) GetByUserName(ctx context.Context, userName string, c *grpc.ClientConn) (usr.User, *xErrors.Error) {
	cnn := userGrpc.NewUserServiceClient(c)
	e, stat := cnn.GetByUserName(context.Background(), &userGrpc.UserName{Username: userName})
	st, ok := status.FromError(stat)
	fmt.Println(e, st, ok)
	return usr.User{
		FirstName:   e.Usr.FirstName,
		LastName:    e.Usr.LastName,
		Email:       e.Usr.Email,
		PhoneNumber: e.Usr.PhoneNumber,
		UserName:    e.Usr.UserName,
		Password:    e.Usr.Password,
	}, nil
}

func (s *Service) Update(ctx context.Context, u *usr.User, c *grpc.ClientConn) *xErrors.Error {
	cnn := userGrpc.NewUserServiceClient(c)
	e, stat := cnn.Create(context.Background(), &userGrpc.UserRequest{
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
		UserName:    u.UserName,
		Password:    u.Password,
	})
	st, ok := status.FromError(stat)
	fmt.Println(e, st, ok)
	return nil
}

func (s *Service) Remove(ctx context.Context, userName string, c *grpc.ClientConn) *xErrors.Error {
	cnn := userGrpc.NewUserServiceClient(c)
	e, stat := cnn.GetByUserName(context.Background(), &userGrpc.UserName{Username: userName})
	st, ok := status.FromError(stat)
	fmt.Println(e, st, ok)
	return nil
}
