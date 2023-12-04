package event

import (
	"fish/client/command"
	"strconv"
)

func ClientConnect(ctx *Context, data string) {
	command.PrintNotice("Connection to server Successful, welcome to poker !! ")
	ctx.UserId, _ = strconv.Atoi(data)
}
