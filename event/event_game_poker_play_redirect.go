package event

import (
	"encoding/json"
	"fish/client/command"
	"fmt"
	"strconv"
)

// var choose = [...]string{"UP", "DOWN"}
var format = "\n[%-4s] %-" + strconv.Itoa(NICKNAME_MAX_LENGTH) + "s  surplus %-2s [%-8s]"

func GamePokerPlayRedirect(ctx *EContext, data string) {
	dataMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(data), &dataMap)

	sellClientId := dataMap["sellClientId"].(string)

	clientInfos := make([]map[string]interface{}, 0)
	clientInfoBytes, _ := json.Marshal(dataMap["clientInfos"])
	_ = json.Unmarshal(clientInfoBytes, &clientInfos)

	for _, clientInfo := range clientInfos {
		position := strconv.Itoa(int(clientInfo["position"].(float64)))
		command.PrintNotice(fmt.Sprintf(format, position, clientInfo["clientNickname"].(string), strconv.Itoa(int(clientInfo["surplus"].(float64))), clientInfo["type"].(string)))
	}
	command.PrintNotice("")
	if sellClientId == ctx.UserId {
		GamePokerPlay(ctx, data)
	} else {
		command.PrintNotice("Next player is " + dataMap["sellClientNickname"].(string) + ". Please wait for him to play his pokers.")
	}
}
