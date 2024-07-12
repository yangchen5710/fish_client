package event

import (
	"fish/client/command"
	"fish/client/common"
	"os"
	"strconv"
	"strings"
)

func ShowOptions(ctx *EContext, data string) {
	command.PrintNotice("Options: ")
	command.PrintNotice("1. PvP")
	command.PrintNotice("2. PvE")
	command.PrintNotice("3. Setting")
	command.PrintNotice("Please enter the number of options (enter [EXIT] log out)")
	command.Write1("options")
	AsynWrite("PushShowOptions", data)
}

func PushShowOptions(ctx *EContext, input *common.Input) {
	line := strings.ToUpper(command.DeletePreAndSufSpace(input.Option))
	if line == "EXIT" {
		os.Exit(0)
	} else {
		choose, err := strconv.Atoi(line)
		if err != nil {
			choose = -1
		}
		switch choose {
		case 1:
			ShowOptionsPVP(ctx, input.Data)
		case 2, 3:
			//ShowOptionsPVE(ctx, input.Data)
			command.PrintNotice("this option is currently not supported, please choose again：")
			ShowOptions(ctx, input.Data)
		//case 3:
		//ShowOptionsSettings(ctx, input.Data)
		default:
			command.PrintNotice("Invalid option, please choose again：")
			ShowOptions(ctx, input.Data)
		}
	}
}
