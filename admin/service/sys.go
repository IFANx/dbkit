package service

import (
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
	"net/http"
)

func GetAllSysInfo(ctx *gin.Context) {
	hostInfo, _ := host.Info()
	memoryInfo, _ := mem.VirtualMemory()
	swapInfo, _ := mem.SwapMemory()
	cpuInfo, _ := cpu.Info()
	diskInfo, _ := disk.Usage(".")
	processIds, _ := process.Pids()

	ctx.JSON(http.StatusOK, gin.H{
		"ok": true,
		"data": map[string]interface{}{
			"host":    hostInfo,
			"memory":  memoryInfo,
			"swap":    swapInfo,
			"cpu":     cpuInfo,
			"process": len(processIds),
			"disk":    diskInfo,
		},
	})
}
