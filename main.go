package main

import (
	"github.com/gin-gonic/gin"
	"main/pkg/conf"
	"main/pkg/logger"
	"main/router"
	"main/service"
)

//	@title			Dorm-changing Backend API
//	@description	Backend system of Dorm-changing mini program, CCNU
//	@version		0.2
//	@contact.name	@a48zhang & @LogSingleDog
//	@contact.email	3557695455@qq.com 1034028483@qq.com
//	@schemes		http
//	@BasePath		/api/v1

func main() {
	gin.SetMode(conf.GetConf("BD_MODE"))

	service.Ping()
	err := router.Register(gin.Default()).Run(":8080")
	logger.FatalLog("Exit. Reason: " + err.Error())
}
