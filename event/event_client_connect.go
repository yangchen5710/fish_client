package event

import (
	"fish/client/command"
)

func ClientConnect(ctx *EContext, data string) {
	command.PrintNotice("Connection to server Successful, welcome to poker !! ")
	ctx.UserId = data
}
