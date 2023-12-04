package main

import (
	"fish/client/common"
	"fish/client/event"
	"fish/client/service"
	"fmt"
	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"log"
	"os"
	"os/signal"
	"strconv"
)

var c *gosocketio.Client

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./executable <IP> <port>")
		return
	}

	ip := os.Args[1]
	port, _ := strconv.Atoi(os.Args[2])

	client, err := gosocketio.Dial(
		gosocketio.GetUrl(ip, port, false),
		transport.GetDefaultWebsocketTransport(),
	)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer client.Close()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	clientChan := make(chan common.Message)
	serverChan := make(chan common.Message)

	ctx := event.NewEventContext(&clientChan, &serverChan)
	ctx.Listen()

	io := service.NewSocketIo(client, &clientChan, &serverChan)
	io.On()
	io.Emit()

	select {}
	/*for {
		select {
		case <-interrupt:
			fmt.Println("Interrupted by user.")
			return
		default:
			time.Sleep(100 * time.Millisecond) // 等待命令执行和服务器响应
		}
	}*/
}
