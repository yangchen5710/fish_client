package event

import (
	"encoding/json"
	"fish/client/command"
	"strconv"
)

func RoomJoinFailByInExist(ctx *EContext, data string) {
	dataMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(data), &dataMap)
	command.PrintNotice("Join room failed. Room " + strconv.Itoa(int(dataMap["roomId"].(float64))) + " inexists!")
	ShowOptions(ctx, data)
}
