package event

import (
	"fish/client/common"
	"fmt"
)

type Room common.Room

type EContext struct {
	clientChan *chan common.Message
	serverChan *chan common.Message
	UserId     int
}

func NewEventContext(clientChan, serverChan *chan common.Message) *EContext {
	return &EContext{
		clientChan: clientChan,
		serverChan: serverChan,
		UserId:     0,
	}
}

func (ctx *EContext) Listen() {
	go func() {
		for {
			transferData := <-*ctx.clientChan
			ctx.call(transferData.Code, transferData.Data)
		}
	}()
}

func (ctx *EContext) call(code string, data string) {
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

func (ctx *EContext) pushToServer(serverCode string, data string) {
	*ctx.serverChan <- common.Message{
		Code: serverCode,
		Data: data,
	}
}

func (ctx *EContext) InitLastSellInfo() {
	//ctx.LastPokers = nil
	//ctx.LastSellClientNickname = ""
	//ctx.LastSellClientType = ""
}
