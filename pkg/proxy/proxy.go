package proxy

import (
	"context"
	"gateway/pkg/service/user"
	"github.com/mhthrh/common_pkg/pkg/logger"
	uu "github.com/mhthrh/common_pkg/pkg/model/user"
	"github.com/mhthrh/common_pkg/pkg/pool/grpcPool"
	"github.com/mhthrh/common_pkg/pkg/xErrors"
	"go.uber.org/zap"
	"time"
)

var (
	pool *grpcPool.GrpcPool
	keys = []string{"userPool"}
)

type Proxy struct {
	L   logger.ILogger
	srv user.Service
}

func New(logger logger.ILogger, p *grpcPool.GrpcPool) (*Proxy, error) {
	pool = p
	return &Proxy{L: logger, srv: *user.New(logger)}, nil
}
func (p *Proxy) Create(ctx context.Context, request uu.User) *xErrors.Error {
	defer func(begin time.Time) {
		p.L.Info(ctx, "Create",
			zap.Duration("took", time.Since(begin)),
		)
	}(time.Now())
	p.L.Info(ctx, "start CreateUser ", zap.Any("inp", request))
	c, err := pool.Get()
	if err != nil {

	}

	return p.srv.Create(ctx, &request, c)
}
func (p *Proxy) Get(ctx context.Context, userName string) (uu.User, *xErrors.Error) {
	defer func(begin time.Time) {
		p.L.Info(ctx, "Get",
			zap.Duration("took", time.Since(begin)),
		)
	}(time.Now())
	p.L.Info(ctx, "start Get user ", zap.String("user name", userName))

	c, err := pool.Get()
	if err != nil {

	}
	return p.srv.GetByUserName(ctx, userName, c)
}
func (p *Proxy) Update(ctx context.Context, request uu.User) *xErrors.Error {
	defer func(begin time.Time) {
		p.L.Info(ctx, "Update",
			zap.Duration("took", time.Since(begin)),
		)
	}(time.Now())
	p.L.Info(ctx, "start Update ", zap.Any("user name", request))
	c, err := pool.Get()
	if err != nil {

	}
	return p.srv.Update(ctx, &request, c)
}
func (p *Proxy) Remove(ctx context.Context, userName string) *xErrors.Error {
	defer func(begin time.Time) {
		p.L.Info(ctx, "Remove",
			zap.Duration("took", time.Since(begin)),
		)
	}(time.Now())
	p.L.Info(ctx, "start Remove ", zap.String("user name", userName))
	c, err := pool.Get()
	if err != nil {

	}
	return p.srv.Remove(ctx, userName, c)
}
