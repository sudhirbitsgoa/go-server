package system

import (
	redis "server/cache"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

/**
  This will read system memory
**/
func readcpu() {
	v, _ := cpu.Percent(0, false)
	redis.Client.LPush("cpuusage", v[0])
}

func CPUIterator() {
	for {
		time.Sleep(time.Second)
		readcpu()
	}
}
