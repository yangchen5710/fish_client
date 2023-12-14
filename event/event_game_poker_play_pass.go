package event

import (
	"encoding/json"
	"fish/client/command"
)

func GamePokerPlayPass(ctx *EContext, data string) {
	dataMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(data), &dataMap)

	command.PrintNotice(dataMap["clientNickname"].(string) + " passed. It is now " + dataMap["nextClientNickname"].(string) + "'s turn.")

	turnClientId := dataMap["nextClientId"].(string)
	if ctx.UserId == turnClientId {
		ctx.pushToServer(SERVER_CODE_GAME_POKER_PLAY_REDIRECT, "")
	}
}
