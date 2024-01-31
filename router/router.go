package router

import (
	"github.com/gin-gonic/gin"
	"main/router/middleware"
)

func Register(e *gin.Engine) *gin.Engine {

	e.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	e.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{
			"message": "Incorrect API route.",
		})
	})

	v1 := e.Group("/api/v1")
	{
		v1.POST("/login")
		v1.Use(middleware.TokenParser)
		users := v1.Group("/users")
		{
			users.GET("/", func(ctx *gin.Context) {
				ctx.JSON(200, gin.H{
					"message": "todo",
				})
			})
		}
	}

	return e
}
