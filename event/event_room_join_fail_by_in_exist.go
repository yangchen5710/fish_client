package event

import (
	"encoding/json"
	"fish/client/command"
)

func RoomJoinFailByInExist(ctx *EContext, data string) {
	dataMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(data), &dataMap)
	command.PrintNotice("Join room failed. Room " + dataMap["roomId"].(string) + " inexists!")
	ShowOptions(ctx, data)
}
