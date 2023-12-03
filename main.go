package main

import (
	"bufio"
	"fmt"
	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"
)

type Message struct {
	Cmd     []byte `json:"cmd"`
	Message string `json:"message"`
}

type Action struct {
	NickName int
	Options  bool
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./executable <IP> <port>")
		return
	}

	ip := os.Args[1]
	port, _ := strconv.Atoi(os.Args[2])

	//reader := bufio.NewReader(os.Stdin)

	input := bufio.NewScanner(os.Stdin)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	c, err := gosocketio.Dial(
		gosocketio.GetUrl(ip, port, false),
		transport.GetDefaultWebsocketTransport(),
	)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	clientChan := make(chan string)

	go func() {
		//defer close(done)
		//for {
		_ = c.On("connectCallback", func(h *gosocketio.Channel, args Message) {
			fmt.Println("Connection to server Successful, welcome to poker !! ")
			clientChan <- "nickname"
		})
		_ = c.On("nickNameCallback", func(h *gosocketio.Channel, args Message) {
			//fmt.Println("testCallback")
			fmt.Println(args.Message)
		})

		_ = c.On("testCallback", func(h *gosocketio.Channel, args Message) {
			//fmt.Println("testCallback")
			fmt.Println(string(args.Cmd))
		})

		_ = c.On("joinCallback", func(h *gosocketio.Channel, args Message) {
			fmt.Printf("Received: %s\n", args)
		})
		//}
	}()
	go func() {
		for {
			transferData := <-clientChan
			fmt.Println(transferData)

		}
	}()

	for {
		select {
		case <-interrupt:
			fmt.Println("Interrupted by user.")
			return
		default:
			input.Scan()
			fmt.Println("text: " + input.Text())
			/*cmd, _ := reader.ReadString('\n')
			cmd = strings.TrimSpace(cmd)

			//err := c.Emit("test", []byte(cmd))
			err := c.Emit("cmd", Message{
				Cmd: []byte(cmd),
			})
			if err != nil {
				fmt.Println("write:", err)
				return
			}*/

			time.Sleep(100 * time.Millisecond) // 等待命令执行和服务器响应
		}
	}
}
