package event

import (
	"fish/client/command"
	"strconv"
	"strings"
)

func ShowOptionsSettings(ctx *EContext, data string) {
	command.PrintNotice("Setting: ")
	command.PrintNotice("1. Card with shape edges (Default)")
	command.PrintNotice("2. Card with rounded edges")
	command.PrintNotice("3. Text Only with types")
	command.PrintNotice("4. Text Only without types")
	command.PrintNotice("5. Unicode Cards")

	command.PrintNotice("Please enter the number of setting (enter [BACK] return options list)")
	line := command.DeletePreAndSufSpace(command.Write("setting"))
	if strings.ToUpper(line) == "BACK" {
		ShowOptions(ctx, data)
	} else {
		choose, e := strconv.Atoi(line)
		if e != nil {
			choose = -1
		}
		if choose >= 1 && choose <= 5 {
			//ctx.PokerPrinterType = choose - 1
			ShowOptions(ctx, data)
		} else {
			command.PrintNotice("Invalid setting, please choose again：")
			ShowOptionsSettings(ctx, data)
		}
	}
}
