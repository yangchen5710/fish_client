package event

import (
	"fish/client/command"
	"strconv"
	"strings"
)

func ShowOptionsPVE(ctx *EContext, data string) {
	command.PrintNotice("PVE: ")
	command.PrintNotice("1. Simple Model")
	command.PrintNotice("2. Medium Model")
	command.PrintNotice("3. Difficulty Model")
	command.PrintNotice("Please enter the number of options (enter [BACK] return options list)")

	line := command.Write("pve")
	if strings.ToUpper(line) == "BACK" {
		ShowOptions(ctx, data)
	} else {
		choose, e := strconv.Atoi(line)
		if e != nil {
			choose = -1
		}
		if choose > 0 && choose < 4 {
			//ctx.InitLastSellInfo()
			ctx.pushToServer(SERVER_CODE_ROOM_CREATE_PVE, strconv.Itoa(choose))
		} else {
			command.PrintNotice("Invalid option, please choose again：")
			ShowOptionsPVE(ctx, data)
		}
	}
}
