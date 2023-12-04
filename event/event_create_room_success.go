package event

import (
	"encoding/json"
	"fish/client/command"
	"strconv"
)

func RoomCreateSuccess(ctx *EContext, data string) {
	room := Room{}
	_ = json.Unmarshal([]byte(data), &room)

	//ctx.InitLastSellInfo()

	command.PrintNotice("You have created a room with id " + strconv.Itoa(room.RoomId))
	command.PrintNotice("Please wait for other players to join !")
}
