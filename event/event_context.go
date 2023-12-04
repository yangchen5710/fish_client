package event

import (
	"fish/client/common"
	"fmt"
)

type Room common.Room

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
	case CLIENT_CONNECT:
		ClientConnect(ctx, data)
	case CLIENT_NICKNAME_SET:
		ClientNicknameSet(ctx, data)
	case SHOW_OPTIONS:
		ShowOptions(ctx, data)
	case ROOM_CREATE_SUCCESS:
		RoomCreateSuccess(ctx, data)
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

func (ctx *Context) InitLastSellInfo() {
	//ctx.LastPokers = nil
	//ctx.LastSellClientNickname = ""
	//ctx.LastSellClientType = ""
}
