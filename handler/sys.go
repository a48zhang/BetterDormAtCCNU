package handler

import (
	"github.com/gin-gonic/gin"
	"main/service"
)
import "github.com/shirou/gopsutil/cpu"

func ServiceStatus(ctx *gin.Context) {
	res, _ := cpu.Info()
	Response(ctx, 200, res, service.GetStatus())

}
