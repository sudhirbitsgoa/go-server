package system

import (
	"fmt"
	redis "go-server/cache"
	"strconv"
	"time"

	"golang.org/x/net/websocket"

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

type memMessage struct {
	Percentage float64   `json:"percentage"`
	Metric     string    `json:"metric"`
	TimeStamp  time.Time `json:"timeStamp"`
}

func SendMemoryWb(ws *websocket.Conn) {
	var err error
	for {
		time.Sleep(time.Second)
		per, error := redis.Client.RPop("memoryusage").Result()
		if error != nil {
			fmt.Print("some error")
		} else {
			per, er := strconv.ParseFloat(per, 64)
			value := memMessage{
				Percentage: per,
				Metric:     "memoryusage",
				TimeStamp:  time.Now(),
			}
			if er == nil {
				// fmt.Println(value)
				if err = websocket.JSON.Send(ws, value); err != nil {
					fmt.Println("Can't send", err)
					break
				}
			}
		}
	}
}
