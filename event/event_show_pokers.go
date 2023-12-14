package event

import (
	"encoding/json"
	"fish/client/command"
)

func ShowPokers(ctx *EContext, data string) {
	dataMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(data), &dataMap)

	//ctx.LastSellClientNickname = dataMap["clientNickname"].(string)
	//ctx.LastSellClientType = dataMap["clientType"].(string)

	command.PrintNotice(dataMap["clientNickname"].(string) + "[" + dataMap["clientType"].(string) + "] played:")

	pokers := make([]Poker, 0)
	pokersBytes, _ := json.Marshal(dataMap["pokers"])
	_ = json.Unmarshal(pokersBytes, &pokers)
	command.PrintPokers(pokers, ctx.PokerPrinterType)

	if dataMap["sellClientNickname"] != nil {
		command.PrintNotice("Next player is " + dataMap["sellClientNickname"].(string) + ". Please wait for him to play his pokers.")
	}
}
