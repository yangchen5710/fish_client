package event

import "fish/client/command"

func ClientNicknameSet(ctx *EContext, data string) {
	command.PrintNotice("please set your nickname")
	nickName := command.Write("nickname")
	ctx.pushToServer(SERVER_CODE_CLIENT_NICKNAME_SET, nickName)

}
