package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"main/pkg/conf"
	"main/router/middleware"
)

func ServiceStatus(ctx *gin.Context) {
	if conf.GetConf("bd_mode") != "debug" {
		ResponseError(ctx, 403, "forbidden")
		return
	}
	cpuinfo, _ := cpu.Info()
	memory, _ := mem.VirtualMemory()
	usage, _ := disk.Usage("/")

	ResponseOK(ctx, gin.H{
		"cpu":    cpuinfo,
		"memory": memory,
		"disk":   usage,
		"static": gin.H{
			"from":  middleware.CountFrom,
			"count": middleware.RequestCount,
		},
	})

}
