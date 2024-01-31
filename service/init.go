package service

import (
	"main/pkg/logger"
)

func init() {
	// model.ConnectMongo()
}

func Ping() {
	logger.InfoLog("Service running.")
	return
}
