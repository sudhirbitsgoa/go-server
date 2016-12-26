package system

import (
	redis "go-server/cache"
	"time"

	"github.com/shirou/gopsutil/mem"
)

/**
  This will read system memory
**/
func readmemory() {
	v, _ := mem.VirtualMemory()
	redis.Client.LPush("memoryusage", v.UsedPercent)
}

func Iterator() {
	for {
		time.Sleep(time.Second)
		readmemory()
	}
}
