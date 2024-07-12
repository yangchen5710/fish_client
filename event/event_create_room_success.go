package event

import (
	"encoding/json"
	"fish/client/command"
	"fish/client/common"
	"strconv"
	"strings"
)

func RoomCreateSuccess(ctx *EContext, data string) {
	room := Room{}
	_ = json.Unmarshal([]byte(data), &room)

	command.PrintNotice("You have created a room with id " + strconv.Itoa(room.RoomId))
	command.PrintNotice("Please wait for other players to join ! Notice: If wait too long, Please enter [LEAVE] to leave out room !")
	AsynWrite("PushRoomLeave", data)
}

func PushRoomLeave(ctx *EContext, input *common.Input) {
	line := strings.ToUpper(command.DeletePreAndSufSpace(input.Option))
	if line == "LEAVE" {
		ctx.pushToServer(SERVER_CODE_ROOM_LEAVE, "")
	}
}
