package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/shirou/gopsutil/mem"
	"main/dao"
	"main/pkg/conf"
	"runtime"
	"time"
)

func init() {
	dao.ConnectMongo()
	nid := conf.GetConf("bd_node_id")
	uuid.SetNodeID([]byte(nid))
	go func() {
		for {
			CurrentStatus = GetStatus()
			time.Sleep(5 * time.Second)
		}
	}()
}

type StatusInfo struct {
	OS         string
	ARCH       string
	MemoryUsed float64
}

var CurrentStatus StatusInfo

func (s StatusInfo) String() string {
	res := "System status: \n"
	res += fmt.Sprintf("OS&ARCH: %s %s \n", s.OS, s.ARCH)
	res += fmt.Sprintf("Memory used: %f  \n", s.MemoryUsed)
	return res
}

func GetStatus() StatusInfo {
	CurrentStatus = StatusInfo{
		OS:   runtime.GOOS,
		ARCH: runtime.GOARCH,
	}
	m, _ := mem.VirtualMemory()
	percent := m.UsedPercent
	CurrentStatus.MemoryUsed = percent
	return CurrentStatus
}
