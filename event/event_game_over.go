package event

import (
	"encoding/json"
	"fish/client/command"
)

func GameOver(ctx *EContext, data string) {
	dataMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(data), &dataMap)

	command.PrintNotice("\nPlayer " + dataMap["winnerNickname"].(string) + "[" + dataMap["winnerType"].(string) + "]" + " won the game")
	command.PrintNotice("Game over, friendship first, competition second\n")
}
