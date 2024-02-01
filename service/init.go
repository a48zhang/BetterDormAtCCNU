package service

import (
	"main/model"
	"main/pkg/logger"
)

func init() {
	model.ConnectMongo()
}

func Ping() {
	logger.InfoLog("Service running.")
	return
}
