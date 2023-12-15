package event

import (
	"encoding/json"
	"fish/client/command"
)

func GameOver(ctx *EContext, data string) {
	dataMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(data), &dataMap)

	command.PrintNotice("\nPlayer " + dataMap["winnerNickname"].(string) + "[" + dataMap["winnerType"].(string) + "]" + " won the game")

	command.PrintNotice("")

	clientInfos := make([]map[string]interface{}, 0)
	clientInfoBytes, _ := json.Marshal(dataMap["clientInfos"])
	_ = json.Unmarshal(clientInfoBytes, &clientInfos)
	for _, clientInfo := range clientInfos {
		pokers := make([]Poker, 0)
		pokersBytes, _ := json.Marshal(clientInfo["pokers"])
		_ = json.Unmarshal(pokersBytes, &pokers)
		command.PrintNotice(clientInfo["clientNickname"].(string) + "[" + clientInfo["type"].(string) + "] remain:")
		command.PrintPokers(pokers, ctx.PokerPrinterType)
	}
	command.PrintNotice("")

	command.PrintNotice("Game over, friendship first, competition second\n")
}
