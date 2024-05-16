package service

import (
	"github.com/shirou/gopsutil/mem"
	"main/model"
	"main/pkg/logger"
	"runtime"
)

func init() {
	model.ConnectMongo()
}

func Ping() {
	logger.InfoLog("Service running.")
	return
}

func GetStatus() string {
	result := "OS&ARCH: " + runtime.GOOS + " " + runtime.GOARCH + "\n"
	m, _ := mem.VirtualMemory()
	return result + m.String()
}
