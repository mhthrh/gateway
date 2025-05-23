package user

import (
	"context"
	"gateway/pkg/proxy"
	"github.com/gin-gonic/gin"
	"github.com/mhthrh/common_pkg/pkg/logger"
	"github.com/mhthrh/common_pkg/pkg/model/user"
	"github.com/mhthrh/common_pkg/pkg/xErrors"
)

var (
	p *proxy.Proxy
)

func New(l logger.ILogger, address string, count int) (err error) {
	p, err = proxy.New(l, address, count)
	return
}
func Create(c *gin.Context) {
	var (
		e *xErrors.Error
		u user.User
	)
	defer func() {
		if e.Code == xErrors.SuccessCode {
			c.JSON(xErrors.GetHttpStatus(e, c.Request.Method), u)
			return
		}
		c.JSON(xErrors.GetHttpStatus(e, c.Request.Method), e)
	}()
	if err := c.ShouldBindJSON(&u); err != nil {
		e = xErrors.FailedResource(nil, nil)
		return
	}

	ctx := context.Background()
	e = p.Create(ctx, u)
}
func GetUser(c *gin.Context) {
	var (
		e   *xErrors.Error
		key = "username"
		u   user.User
	)
	defer func() {
		if e.Code == xErrors.SuccessCode {
			c.JSON(xErrors.GetHttpStatus(e, c.Request.Method), u)
			return
		}
		c.JSON(xErrors.GetHttpStatus(e, c.Request.Method), e)
	}()

	userName, ok := c.GetQuery(key)
	if !ok || userName == "" {
		e = xErrors.NewErrKeyNotExist(key)
		return
	}
	ctx := context.Background()
	u, e = p.Get(ctx, "")
}
func UpdateUser(c *gin.Context) {
	var (
		e *xErrors.Error
		u user.User
	)

	defer func() {
		if e.Code == xErrors.SuccessCode {
			c.JSON(xErrors.GetHttpStatus(e, c.Request.Method), u)
			return
		}
		c.JSON(xErrors.GetHttpStatus(e, c.Request.Method), e)
	}()
	if err := c.ShouldBindJSON(&u); err != nil {
		e = xErrors.FailedResource(nil, nil)
		return
	}

	ctx := context.Background()
	e = p.Update(ctx, u)
}
func DeleteUser(c *gin.Context) {
	var (
		e   *xErrors.Error
		key = "username"
		u   user.User
	)

	defer func() {
		if e.Code == xErrors.SuccessCode {
			c.JSON(xErrors.GetHttpStatus(e, c.Request.Method), u)
			return
		}
		c.JSON(xErrors.GetHttpStatus(e, c.Request.Method), e)
	}()
	userName, ok := c.GetQuery(key)
	if !ok || userName == "" {
		e = xErrors.NewErrKeyNotExist(key)
		return
	}
	ctx := context.Background()
	e = p.Remove(ctx, userName)
}
