package event

import (
	"encoding/json"
	"fish/client/command"
	"fmt"
	"strconv"
)

func ShowRooms(ctx *EContext, data string) {
	roomList := make(map[int]*Room)
	_ = json.Unmarshal([]byte(data), &roomList)

	//fmt.Println(roomList)
	if len(roomList) > 0 {
		format := "#\t%s\t|\t%-" + strconv.Itoa(NICKNAME_MAX_LENGTH) + "s\t|\t%-6s\t|\t%-6s\t#"
		command.PrintNotice(fmt.Sprintf(format, "ID", "OWNER", "COUNT", "TYPE"))
		for _, room := range roomList {
			command.PrintNotice(fmt.Sprintf(format, strconv.Itoa(room.RoomId), room.RoomOwner, strconv.Itoa(room.RoomClientCount), strconv.Itoa(room.RoomType)))
		}
		command.PrintNotice("")
		ShowOptionsPVP(ctx, data)
	} else {
		command.PrintNotice("No available room, please create a room ！")
		ShowOptionsPVP(ctx, data)
	}

}
