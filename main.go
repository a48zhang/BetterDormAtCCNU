package main

import (
	"github.com/gin-gonic/gin"
	"main/pkg/conf"
	"main/pkg/logger"
	"main/router"
	"main/service"
)

func main() {
	gin.SetMode(conf.GetConf("BD_MODE"))

	service.Ping()
	err := router.Register(gin.Default()).Run(":8080")
	logger.FatalLog("Exit. Reason: " + err.Error())
}
