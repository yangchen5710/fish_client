package event

import (
	"fish/client/command"
	"fish/client/common"
	"fmt"
)

type Room common.Room
type Poker = common.Poker

const NICKNAME_MAX_LENGTH = 10

var Input = new(common.Input)

type EContext struct {
	clientChan       *chan common.Message
	serverChan       *chan common.Message
	inputChan        *chan *common.Input
	UserId           string
	PokerPrinterType int
}

func NewEventContext(clientChan, serverChan *chan common.Message, inputChan *chan *common.Input) *EContext {
	return &EContext{
		clientChan:       clientChan,
		serverChan:       serverChan,
		inputChan:        inputChan,
		UserId:           "",
		PokerPrinterType: 0,
	}
}

func (ctx *EContext) Listen() {
	go func() {
		for {
			transferData := <-*ctx.clientChan
			ctx.call(transferData.Code, transferData.Data)
		}
	}()

	go func() {
		for command.Input.Scan() {
			Input.Option = command.Input.Text()
			*ctx.inputChan <- Input
		}
	}()

	go func() {
		for {
			input := <-*ctx.inputChan
			ctx.callInput(input)
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
	case SHOW_ROOMS:
		ShowRooms(ctx, data)
	case ROOM_JOIN_SUCCESS:
		RoomJoinSuccess(ctx, data)
	case GAME_STARTING:
		GameStarting(ctx, data)
	case GAME_LANDLORD_ELECT:
		GameLandlordElect(ctx, data)
	case ROOM_JOIN_FAIL_BY_FULL:
		RoomJoinFailByFull(ctx, data)
	case ROOM_JOIN_FAIL_BY_INEXIST:
		RoomJoinFailByInExist(ctx, data)
	case GAME_LANDLORD_CYCLE:
		GameLandlordCycle(ctx, data)
	case GAME_LANDLORD_CONFIRM:
		GameLandlordConfirm(ctx, data)
	case GAME_POKER_PLAY_REDIRECT:
		GamePokerPlayRedirect(ctx, data)
	case GAME_POKER_PLAY_MISMATCH:
		GamePokerPlayMismatch(ctx, data)
	case SHOW_POKERS:
		ShowPokers(ctx, data)
	case GAME_POKER_PLAY_PASS:
		GamePokerPlayPass(ctx, data)
	case GAME_POKER_PLAY_CANT_PASS:
		GamePokerPlayCantPass(ctx, data)
	case GAME_OVER:
		GameOver(ctx, data)
	case CLIENT_EXIT:
		ClientExit(ctx, data)
	case ROOM_OWNER_SELECT:
		RoomOwnerSelect(ctx, data)
	case "ROOM_LEAVE_SUCCESS":
		RoomLeaveSuccess(ctx, data)
	default:
		fmt.Println("undefined Code")
	}
}

func (ctx *EContext) pushToServer(serverCode string, data string) {
	*ctx.serverChan <- common.Message{
		Code: serverCode,
		Data: data,
	}
}

func (ctx *EContext) InitLastSellInfo() {

}

func AsynWrite(funcName, data string) {
	Input.FuncName = funcName
	Input.Data = data
}

func CleanInput(input *common.Input) {
	Input.FuncName = ""
	Input.Data = ""
	Input.Option = ""
}

func (ctx *EContext) callInput(input *common.Input) {
	//fmt.Println(input)
	switch input.FuncName {
	case "PushNicknameSet":
		PushNicknameSet(ctx, input)
	case "PushShowOptions":
		PushShowOptions(ctx, input)
	case "PushShowOptionsPVP":
		PushShowOptionsPVP(ctx, input)
	case "PushRoomSelect":
		PushRoomSelect(ctx, input)
	case "PushRoomLeave":
		PushRoomLeave(ctx, input)
	case "PushGameLandlordElect":
		PushGameLandlordElect(ctx, input)
	case "PushGamePokerPlay":
		PushGamePokerPlay(ctx, input)
	case "PushRoomOwnerSelect":
		PushRoomOwnerSelect(ctx, input)
	case "":
	default:
		fmt.Println("undefined Code")
	}
}
