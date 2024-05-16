package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "main/docs"
	"main/handler"
	"main/router/middleware"
)

func Register(e *gin.Engine) *gin.Engine {

	e.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{
			"message": "Incorrect API route.",
		})
	})

	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	e.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong"})
	})

	v1 := e.Group("/api/v1")
	{
		v1.POST("register", handler.Register)
		v1.POST("/login", handler.Login)

		internal := v1.Group("/internal")
		internal.GET("/pdf", handler.GenPDFCallback)

		v1.Use(middleware.TokenParser)

		users := v1.Group("/users")
		forms := v1.Group("/forms")
		reports := v1.Group("/report")

		users.GET("/", handler.GetUserInfo)
		users.POST("/", handler.UpdateMyInfo)

		forms.GET("/my", handler.GetForms)
		forms.POST("/create", handler.CreateForm)
		forms.GET("/", handler.GetOneForm)
		forms.POST("/check", handler.CheckForm)
		forms.GET("/assigned", handler.GetAssignedForms)
		forms.POST("/pdf", handler.GenFormPDF)
		forms.GET("/pdf", handler.GenFormByID)

		reports.POST("/", handler.NewReport)
		reports.GET("/", handler.MyReport)

	}

	return e
}
