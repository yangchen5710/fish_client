package event

import (
	"fish/client/command"
	"os"
	"strconv"
	"strings"
)

func ShowOptions(ctx *Context, data string) {
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
			//showOptionsPvP()
		case 2:
			//showOptionsPvP()
		case 3:
			//showOptionsPvP()
		default:
			command.PrintNotice("Invalid option, please choose again：")
			ShowOptions(ctx, data)
		}

	}
}
