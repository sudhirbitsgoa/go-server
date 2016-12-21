package system

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/mem"
)

/**
  This will read system memory
**/
func readmemory() *mem.VirtualMemoryStat {
	v, _ := mem.VirtualMemory()
	return v
}

func Iterator() {
	for {
		time.Sleep(time.Second)
		fmt.Print(time.Now(), readmemory())
	}
}
