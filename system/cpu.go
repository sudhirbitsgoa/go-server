package system

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

/**
  This will read system memory
**/
func readcpu() []float64 {
	v, _ := cpu.Percent(0, false)

	// almost every return value is a struct
	// fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v., v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	// fmt.Println("this is cool", v)
	return v
}

func CPUIterator() {
	c := time.Tick(time.Second)
	for now := range c {
		fmt.Println(now, readcpu())
	}
}
