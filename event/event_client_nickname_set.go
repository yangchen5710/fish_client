package event

import (
	"fish/client/command"
	"fish/client/common"
	"strings"
)

func ClientNicknameSet(ctx *EContext, data string) {
	command.PrintNotice("please set your nickname")

	command.Write1("nickname")
	AsynWrite("PushNicknameSet", data)
}

func PushNicknameSet(ctx *EContext, input *common.Input) {
	nickname := strings.ToUpper(command.DeletePreAndSufSpace(input.Option))
	if nickname == "" {
		command.PrintNotice("Invalid nickname, please input again：")
		ClientNicknameSet(ctx, input.Data)
		return
	}
	ctx.pushToServer(SERVER_CODE_CLIENT_NICKNAME_SET, nickname)
}
