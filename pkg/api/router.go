package api

import (
	"gateway/pkg/api/user"
	"github.com/gin-gonic/gin"
	"github.com/mhthrh/common_pkg/pkg/xErrors"
	"net/http"
)

func Run() http.Handler {
	g := gin.Default()
	g.Use(checkToken())

	userGroup := g.Group("/user")
	//notificationGroup:=g.Group("/notification")

	userGroup.Use(checkAddress())

	userGroup.POST("/create", user.Create)
	userGroup.GET("/get", user.GetUser)
	userGroup.PUT("/update", user.UpdateUser)
	userGroup.DELETE("/delete", user.DeleteUser)

	g.NoRoute(func(context *gin.Context) {
		context.JSON(xErrors.GetHttpStatus(xErrors.NotImplemented(context.Request.Method), context.Request.Method), xErrors.NotImplemented(context.Request.Method))
	})
	g.NoMethod(func(context *gin.Context) {
		context.JSON(xErrors.GetHttpStatus(xErrors.NotImplemented(context.Request.Method), context.Request.Method), xErrors.NotImplemented(context.Request.Method))
	})
	return g
}
