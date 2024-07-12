package event

import (
	"fish/client/command"
	"fish/client/common"
	"strconv"
	"strings"
)

func ShowOptionsPVP(ctx *EContext, data string) {
	command.PrintNotice("PVP: ")
	command.PrintNotice("1. Create Room")
	command.PrintNotice("2. Room List")
	command.PrintNotice("3. Join Room")
	command.PrintNotice("Please enter the number of options (enter [BACK] return options list)")
	command.Write1("pvp")
	AsynWrite("PushShowOptionsPVP", data)
}

func PushShowOptionsPVP(ctx *EContext, input *common.Input) {
	line := strings.ToUpper(command.DeletePreAndSufSpace(input.Option))
	if line == "BACK" {
		ShowOptions(ctx, input.Data)
	} else {
		choose, e := strconv.Atoi(line)
		if e != nil {
			choose = -1
		}
		switch choose {
		case 1:
			ctx.pushToServer(SERVER_CODE_ROOM_CREATE, "")
		case 2:
			ctx.pushToServer(SERVER_CODE_GET_ROOMS, "")
		case 3:
			command.PrintNotice("Please enter the room id you want to join (enter [BACK] return options list)")
			command.Write1("roomid")
			AsynWrite("PushRoomSelect", input.Data)

		default:
			command.PrintNotice("Invalid option, please choose again：")
			ShowOptionsPVP(ctx, input.Data)
		}
	}
}

func PushRoomSelect(ctx *EContext, input *common.Input) {
	line := command.DeletePreAndSufSpace(input.Option)
	if strings.ToUpper(line) == "BACK" {
		ShowOptionsPVP(ctx, input.Data)
	} else {
		roomId, e := strconv.Atoi(line)
		if e != nil {
			roomId = -1
		}
		if roomId < 1 {
			command.PrintNotice("Invalid options, please choose again：")
			ShowOptionsPVP(ctx, input.Data)
		} else {
			CleanInput(input)
			ctx.pushToServer(SERVER_CODE_ROOM_JOIN, strconv.Itoa(roomId))
		}
	}
}
