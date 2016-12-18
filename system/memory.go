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

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v)
	return v
}

func Iterator() {
	c := time.Tick(time.Second)
	for now := range c {
		fmt.Print(now, readmemory())
	}
}
