package system

import (
	"fmt"
	redis "go-server/cache"
	"net/http"
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

func EchoCPU(w http.ResponseWriter, r *http.Request) {
	// fmt.Print("this is called")
	// c, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	// if err != nil {
	// 	log.Print("upgrade:", err)
	// 	return
	// }
	// defer c.Close()
	// for {
	// 	mt, message, err := c.ReadMessage()
	// 	if err != nil {
	// 		log.Println("read:", err)
	// 		break
	// 	}
	// 	log.Printf("recv: %s", message)
	// 	err = c.WriteMessage(mt, message)
	// 	if err != nil {
	// 		log.Println("write:", err)
	// 		break
	// 	}
	// }
}

type message struct {
	// the json tag means this will serialize as a lowercased field
	TimeStamp time.Time `json:"timeStamp"`
	Perct     float64   `json:"perct"`
}

func EchoWS(ws *websocket.Conn) {
	var err error

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)

		msg := reply
		fmt.Println("Sending to client: " + msg)

		switch msg {
		case "cpu":
			for {
				time.Sleep(time.Second)
				per, error := redis.Client.RPop("cpuusage").Result()
				if error != nil {
					fmt.Print("some error")
				} else {
					per, er := strconv.ParseFloat(per, 64)
					value := message{
						Perct:     per,
						TimeStamp: time.Now(),
					}
					// jsonValue, eerr := json.Marshal(value)
					if er == nil {
						// fmt.Println(value)
						if err = websocket.JSON.Send(ws, value); err != nil {
							fmt.Println("Can't send", err)
							break
						}
					}
				}
			}
		default:
			if err = websocket.Message.Send(ws, msg); err != nil {
				fmt.Println("Can't send")
				break
			}
		}
	}
}
