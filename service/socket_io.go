package service

import (
	"fish/client/common"
	gosocketio "github.com/graarh/golang-socketio"
)

type SocketIo struct {
	client     *gosocketio.Client
	clientChan *chan common.Message
	serverChan *chan common.Message
}

func NewSocketIo(client *gosocketio.Client, clientChan, serverChan *chan common.Message) *SocketIo {
	return &SocketIo{
		client:     client,
		clientChan: clientChan,
		serverChan: serverChan,
	}
}

func (io *SocketIo) On() {
	go func() {
		_ = io.client.On("eventCallback", func(h *gosocketio.Channel, args common.Message) {
			*io.clientChan <- args
		})
	}()
}

func (io *SocketIo) Emit() {
	go func() {
		for {
			transferData := <-*io.serverChan
			_ = io.client.Emit("event", transferData)

		}
	}()
}
