package event

import (
	"encoding/json"
	"fish/client/command"
	"strconv"
	"strings"
)

func RoomCreateSuccess(ctx *EContext, data string) {
	room := Room{}
	_ = json.Unmarshal([]byte(data), &room)

	//ctx.InitLastSellInfo()

	command.PrintNotice("You have created a room with id " + strconv.Itoa(room.RoomId))
	command.PrintNotice("Please wait for other players to join ! Notice: If wait too long, Please enter [LEAVE] to leave out room !")
	line := strings.ToUpper(command.DeletePreAndSufSpace(command.Write("options")))
	if line == "LEAVE" {
		ctx.pushToServer(SERVER_CODE_ROOM_LEAVE, "")
	}
}
