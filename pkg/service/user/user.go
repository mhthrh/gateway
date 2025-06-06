package user

import (
	"context"
	"github.com/mhthrh/common_pkg/pkg/logger"
	usr "github.com/mhthrh/common_pkg/pkg/model/user"
	userGrpc "github.com/mhthrh/common_pkg/pkg/model/user/grpc/v1"
	"github.com/mhthrh/common_pkg/pkg/xErrors"
	"go.uber.org/zap"
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
	s.l.Info(ctx, "start user.creat ", zap.Any("input", u))

	s.l.Info(ctx, "make a connection to grpc server")
	cnn := userGrpc.NewUserServiceClient(c)
	s.l.Info(ctx, "server connect")

	s.l.Info(ctx, "grpc calling,Create ")

	_, gErr := cnn.Create(context.Background(), &userGrpc.UserRequest{
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
		UserName:    u.UserName,
		Password:    u.Password,
	})

	if gErr != nil {
		st, ok := status.FromError(gErr)
		s.l.Error(ctx, "grpc result", zap.Any("stat", gErr), zap.Any("st", st), zap.Any("ok", ok))
		e, _ := xErrors.Yaml2Struct([]byte(st.Message()))
		return e
	}
	s.l.Info(ctx, "finish user.create ")

	return xErrors.Success()
}

func (s *Service) GetByUserName(ctx context.Context, userName string, c *grpc.ClientConn) (usr.User, *xErrors.Error) {
	s.l.Info(ctx, "start user.GetByUserName ", zap.Any("input", userName))

	s.l.Info(ctx, "make a connection to grpc server")
	cnn := userGrpc.NewUserServiceClient(c)
	s.l.Info(ctx, "server connect")

	s.l.Info(ctx, "grpc calling,Create ")
	e, gErr := cnn.GetByUserName(context.Background(), &userGrpc.UserName{Username: userName})
	if gErr != nil {
		st, ok := status.FromError(gErr)
		s.l.Error(ctx, "grpc result", zap.Any("stat", gErr), zap.Any("st", st), zap.Any("ok", ok))
		e, _ := xErrors.Yaml2Struct([]byte(st.Message()))
		return usr.User{}, e
	}
	s.l.Info(ctx, "finish user.getByUserName ")

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
	s.l.Info(ctx, "start user.Update ", zap.Any("input", u))

	s.l.Info(ctx, "make a connection to grpc server")
	cnn := userGrpc.NewUserServiceClient(c)
	s.l.Info(ctx, "server connect")

	s.l.Info(ctx, "grpc calling,Create ")
	_, gErr := cnn.Update(context.Background(), &userGrpc.UserRequest{
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
		UserName:    u.UserName,
		Password:    u.Password,
	})
	if gErr != nil {
		st, ok := status.FromError(gErr)
		s.l.Error(ctx, "grpc result", zap.Any("stat", gErr), zap.Any("st", st), zap.Any("ok", ok))
		e, _ := xErrors.Yaml2Struct([]byte(st.Message()))
		return e
	}
	s.l.Info(ctx, "finish user.update ")

	return xErrors.Success()
}

func (s *Service) Remove(ctx context.Context, userName string, c *grpc.ClientConn) *xErrors.Error {
	s.l.Info(ctx, "start user.Remove ", zap.Any("input", userName))

	s.l.Info(ctx, "make a connection to grpc server")
	cnn := userGrpc.NewUserServiceClient(c)
	s.l.Info(ctx, "server connect")

	s.l.Info(ctx, "grpc calling,Create ")
	_, gErr := cnn.Remove(context.Background(), &userGrpc.UserName{Username: userName})
	if gErr != nil {
		st, ok := status.FromError(gErr)
		s.l.Error(ctx, "grpc result", zap.Any("stat", gErr), zap.Any("st", st), zap.Any("ok", ok))
		e, _ := xErrors.Yaml2Struct([]byte(st.Message()))
		return e
	}
	s.l.Info(ctx, "finish user.Remove ")
	return xErrors.Success()
}
