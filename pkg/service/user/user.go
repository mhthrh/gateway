package user

import (
	"context"
	"fmt"
	"github.com/mhthrh/common_pkg/pkg/logger"
	usr "github.com/mhthrh/common_pkg/pkg/model/user"
	userGrpc "github.com/mhthrh/common_pkg/pkg/model/user/grpc/v1"
	"github.com/mhthrh/common_pkg/pkg/xErrors"
	"google.golang.org/grpc/status"
)

type Service struct {
	uCnn userGrpc.UserServiceClient
	l    logger.ILogger
}

func (s Service) Create(ctx context.Context, user *usr.User) *xErrors.Error {
	e, stat := s.uCnn.Create(context.Background(), &userGrpc.UserRequest{
		FirstName:   "",
		LastName:    "",
		Email:       "",
		PhoneNumber: "",
		UserName:    "",
		Password:    "",
	})
	st, ok := status.FromError(stat)
	fmt.Println(e, st, ok)
	return nil
}

func (s Service) GetByUserName(ctx context.Context, userName string) (usr.User, *xErrors.Error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) Update(ctx context.Context, user *usr.User) *xErrors.Error {
	//TODO implement me
	panic("implement me")
}

func (s Service) Remove(ctx context.Context, userName string) *xErrors.Error {
	//TODO implement me
	panic("implement me")
}
