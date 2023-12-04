package event

import (
	"fish/client/common"
	"fmt"
)

type Context struct {
	clientChan *chan common.Message
	serverChan *chan common.Message
	UserId     int
}

func NewEventContext(clientChan, serverChan *chan common.Message) *Context {
	return &Context{
		clientChan: clientChan,
		serverChan: serverChan,
		UserId:     0,
	}
}

func (ctx *Context) Listen() {
	go func() {
		for {
			transferData := <-*ctx.clientChan
			ctx.call(transferData.Code, transferData.Data)
		}
	}()
}

func (ctx *Context) call(code string, data string) {
	switch code {
	case "clientConnect":
		ClientConnect(ctx, data)
	case "clientNicknameSet":
		ClientNicknameSet(ctx, data)
	case "showOptions":
		ShowOptions(ctx, data)
	default:
		fmt.Println("undefined Code")
		//fmt.Println("choose: " + choose)
		/*_ = c.Emit("serverNicknameSet", Message{
			Message: option,
		})*/
	}
}

func (ctx *Context) pushToServer(serverCode string, data string) {
	*ctx.serverChan <- common.Message{
		Code: serverCode,
		Data: data,
	}
}
