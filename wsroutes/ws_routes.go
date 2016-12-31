package wsroutes

import (
	"fmt"
	sy "go-server/system"

	"golang.org/x/net/websocket"
)

func RoutesHandler(ws *websocket.Conn) {
	var err error
	for {
		var reply string
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("Received back from client: " + reply)

		msg := reply
		fmt.Println("Sending to client: " + msg)

		switch msg {
		case "cpu":
			go sy.SendCPUWb(ws)
		case "memory":
			go sy.SendMemoryWb(ws)
		default:
			if err = websocket.Message.Send(ws, msg); err != nil {
				fmt.Println("Can't send")
				break
			}
		}
	}
}
