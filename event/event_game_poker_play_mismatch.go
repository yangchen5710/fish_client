package event

import (
	"fish/client/command"
)

func GamePokerPlayMismatch(ctx *EContext, data string) {
	/*dataMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(data), &dataMap)

	command.PrintNotice("Your pokers' type is " + dataMap["playType"].(string) + " (" + strconv.Itoa(int(dataMap["playCount"].(float64))) +
		") but previous pokers' type is " + dataMap["preType"].(string) + " (" + strconv.Itoa(int(dataMap["preCount"].(float64))) + "), mismatch !!")*/

	/*if ctx.LastPokers != nil {
		command.PrintNotice(ctx.LastSellClientNickname + "[" + ctx.LastSellClientType + "] played:")
		command.PrintPokers(*ctx.LastPokers, ctx.PokerPrinterType)
	}*/

	command.PrintNotice("Your pokers is mismatch !!")
	ctx.pushToServer(SERVER_CODE_GAME_POKER_PLAY_REDIRECT, "")
}
