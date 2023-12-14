package event

import "fish/client/command"

func ClientExit(ctx *EContext, data string) {
	/*dataMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(data), &dataMap)

	exitClientId := dataMap["exitClientId"].(string)
	var role string
	if exitClientId == ctx.UserId {
		role = "You"
	} else {
		role = dataMap["exitClientNickname"].(string)
	}
	command.PrintNotice(role + " exit from the room. Room disbanded!!\n")*/
	command.PrintNotice(" Room disbanded!!\n")
	ShowOptions(ctx, data)
}
