package system

import (
	"fmt"
	redis "go-server/cache"
	"strconv"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"golang.org/x/net/websocket"
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

type message struct {
	// the json tag means this will serialize as a lowercased field
	TimeStamp time.Time `json:"timeStamp"`
	Metric    string    `json:"metric"`
	Perct     float64   `json:"perct"`
}

func SendCPUWb(ws *websocket.Conn, ch chan bool) {
	var err error
	for {
		select {
		case <-ch:
			return
		default:
			time.Sleep(time.Second)
			per, error := redis.Client.RPop("cpuusage").Result()
			if error != nil {
				fmt.Print("some error")
			} else {
				per, er := strconv.ParseFloat(per, 64)
				value := message{
					Perct:     per,
					Metric:    "cpuusage",
					TimeStamp: time.Now(),
				}
				// jsonValue, eerr := json.Marshal(value)
				if er == nil {
					if err = websocket.JSON.Send(ws, value); err != nil {
						fmt.Println("Can't send", err)
						break
					}
				}
			}
		}
	}
}
