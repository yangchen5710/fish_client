package event

import "fish/client/command"

func ClientNicknameSet(ctx *EContext, data string) {
	command.PrintNotice("please set your nickname")
	nickName := command.Write("nickname")
	if nickName == "" {
		command.PrintNotice("Invalid nickname, please input again：")
		ClientNicknameSet(ctx, data)
		return
	}
	ctx.pushToServer(SERVER_CODE_CLIENT_NICKNAME_SET, nickName)

}
