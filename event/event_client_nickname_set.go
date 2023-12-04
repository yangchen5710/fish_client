package event

import "fish/client/command"

func ClientNicknameSet(ctx *Context, data string) {
	command.PrintNotice("please set your nickname")
	nickName := command.Write("nickname")
	ctx.pushToServer("serverNicknameSet", nickName)

}
