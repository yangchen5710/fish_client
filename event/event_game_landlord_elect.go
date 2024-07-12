package event

import (
	"encoding/json"
	"fish/client/command"
	"fish/client/common"
	"os"
	"strconv"
	"strings"
)

func GameLandlordElect(ctx *EContext, data string) {
	dataMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(data), &dataMap)
	turnClientId := dataMap["nextClientId"].(string)

	/*if dataMap["preClientNickname"] != nil {
		command.PrintNotice(dataMap["preClientNickname"].(string) + " don't rob the landlord!")
	}*/
	if turnClientId == ctx.UserId {
		score := int(dataMap["currentScore"].(float64))
		command.PrintNotice("current score is " + strconv.Itoa(score) + ", please enter score ")
		command.PrintNotice("It's your turn. Do you want to rob the landlord? [1/2/3] (enter [EXIT] to exit current room) (enter [PASS] to pass turn)")
		command.Write1("1/2/3")
		AsynWrite("PushGameLandlordElect", data)
	} else {
		command.PrintNotice("It's " + dataMap["nextClientNickname"].(string) + "'s turn. Please wait patiently for his/her confirmation !")
	}
}

func PushGameLandlordElect(ctx *EContext, input *common.Input) {
	dataMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(input.Data), &dataMap)
	line := strings.ToUpper(command.DeletePreAndSufSpace(input.Option))
	if line == "EXIT" {
		os.Exit(0)
	} else if line == "PASS" {
		CleanInput(input)
		ctx.pushToServer(SERVER_CODE_GAME_LANDLORD_ELECT, "0")
	} else {
		score := int(dataMap["currentScore"].(float64))
		choose, err := strconv.Atoi(line)
		if err != nil || choose <= score {
			choose = -1
		}
		switch choose {
		case 1, 2, 3:
			CleanInput(input)
			ctx.pushToServer(SERVER_CODE_GAME_LANDLORD_ELECT, strconv.Itoa(choose))
		default:
			command.PrintNotice("Invalid options")
			GameLandlordElect(ctx, input.Data)
		}
	}

}
