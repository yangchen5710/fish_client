package event

import (
	"encoding/json"
	"fish/client/command"
	"strconv"
)

func RoomJoinFailByFull(ctx *EContext, data string) {
	dataMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(data), &dataMap)

	command.PrintNotice("Join room failed. Room " + strconv.Itoa(dataMap["roomId"].(int)) + " player count is full!")
	ShowOptions(ctx, data)
}
