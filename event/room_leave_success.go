package event

import (
	"encoding/json"
	"fish/client/command"
	"strconv"
)

func RoomLeaveSuccess(ctx *EContext, data string) {
	dataMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(data), &dataMap)

	ctx.InitLastSellInfo()
	joinClientId, _ := dataMap["clientId"].(string)

	if ctx.UserId == joinClientId {
		command.PrintNotice("You have leave out room：" + strconv.Itoa(int(dataMap["roomId"].(float64))))
		ShowOptions(ctx, data)
		//command.PrintNotice("Please wait for other players to join, start a good game when the room player reaches three !")
	} else {
		command.PrintNotice(dataMap["clientNickname"].(string) + " leave out room, the current number of room player is " + strconv.Itoa(int(dataMap["roomClientCount"].(float64))))
	}
	//os.Exit(0)
}
