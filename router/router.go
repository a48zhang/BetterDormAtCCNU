package router

import (
	"github.com/gin-gonic/gin"
	"main/handler"
	"main/router/middleware"
)

func Register(e *gin.Engine) *gin.Engine {

	e.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{
			"message": "Incorrect API route.",
		})
	})

	// Group test: for developing use, will be removed in the future.
	test := e.Group("/test")
	test.POST("register", handler.Register)
	test.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong"})
	})

	v1 := e.Group("/api/v1")
	{
		v1.POST("/login", handler.Login)
		v1.Use(middleware.TokenParser)
		users := v1.Group("/users")
		users.GET("/", handler.GetUserInfo)

	}

	return e
}
