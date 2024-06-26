package event

import (
	"fish/client/common"
	"fmt"
)

type Room common.Room
type Poker = common.Poker

const NICKNAME_MAX_LENGTH = 10

type EContext struct {
	clientChan       *chan common.Message
	serverChan       *chan common.Message
	UserId           string
	PokerPrinterType int
}

func NewEventContext(clientChan, serverChan *chan common.Message) *EContext {
	return &EContext{
		clientChan:       clientChan,
		serverChan:       serverChan,
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
	//ctx.LastPokers = nil
	//ctx.LastSellClientNickname = ""
	//ctx.LastSellClientType = ""
}
