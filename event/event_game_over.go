package event

import (
	"encoding/json"
	"fish/client/command"
	"fmt"
	"strconv"
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
		if clientInfo["surplus"].(float64) == 0 {
			continue
		}
		pokers := make([]Poker, 0)
		pokersBytes, _ := json.Marshal(clientInfo["pokers"])
		_ = json.Unmarshal(pokersBytes, &pokers)
		command.PrintNotice(clientInfo["clientNickname"].(string) + "[" + clientInfo["type"].(string) + "] remain:")
		command.PrintPokers(pokers, ctx.PokerPrinterType)
	}
	command.PrintNotice("")
	format := "#\t%-" + strconv.Itoa(NICKNAME_MAX_LENGTH) + "s\t|\t%-8s\t|\t%-6s\t#"
	command.PrintNotice(fmt.Sprintf(format, "NICKNAME", "TYPE", "POINT"))
	for _, clientInfo := range clientInfos {
		command.PrintNotice(fmt.Sprintf(format, clientInfo["clientNickname"].(string), clientInfo["type"].(string), strconv.Itoa(int(clientInfo["point"].(float64)))))
	}
	command.PrintNotice("")
	command.PrintNotice("Game over, friendship first, competition second\n")
}
