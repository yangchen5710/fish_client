package event

import (
	"fish/client/command"
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
			ShowOptionsPVP(ctx, data)
		case 2, 3:
			//ShowOptionsPVE(ctx, data)
			command.PrintNotice("this option is currently not supported, please choose again：")
			ShowOptions(ctx, data)
		//case 3:
		//ShowOptionsSettings(ctx, data)
		default:
			command.PrintNotice("Invalid option, please choose again：")
			ShowOptions(ctx, data)
		}
	}
}
