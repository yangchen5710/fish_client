package event

import (
	"fish/client/command"
	"os"
	"strconv"
	"strings"
)

func RoomOwnerSelect(ctx *EContext, data string) {
	command.PrintNotice("Options: ")
	command.PrintNotice("1. RESTARTING GAME")
	command.PrintNotice("2. DISBAND ROOM")
	command.PrintNotice("Please enter the number of options (enter [EXIT] log out)")
	line := strings.ToUpper(command.DeletePreAndSufSpace(command.Write("options")))
	if line == "EXIT" {
		os.Exit(0)
	} else {
		choose, err := strconv.Atoi(line)
		if err != nil {
			choose = -1
		}
		switch choose {
		case 1:
			//ShowOptionsPVP(ctx, data)
			ctx.pushToServer(SERVER_CODE_GAME_STARTING, "")
		case 2:
			//ShowOptionsPVP(ctx, data)
			ctx.pushToServer(SERVER_CODE_ROOM_DISBAND, "")
		default:
			command.PrintNotice("Invalid option, please choose again：")
			RoomOwnerSelect(ctx, data)
		}
	}
}
