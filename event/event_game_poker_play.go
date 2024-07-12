package event

import (
	"encoding/json"
	"fish/client/command"
	"fish/client/common"
	"strings"
)

func GamePokerPlay(ctx *EContext, data string) {
	dataMap := make(map[string]interface{})
	_ = json.Unmarshal([]byte(data), &dataMap)

	command.PrintNotice("It's your turn to play, your pokers are as follows: ")

	pokers := make([]Poker, 0)
	pokersBytes, _ := json.Marshal(dataMap["pokers"])
	_ = json.Unmarshal([]byte(pokersBytes), &pokers)
	command.PrintPokers(pokers, ctx.PokerPrinterType)

	command.PrintNotice("Please enter the card you came up with (enter [EXIT] to exit current room, enter [PASS] to jump current round)")

	command.Write1("card")
	AsynWrite("PushGamePokerPlay", data)
}

func PushGamePokerPlay(ctx *EContext, input *common.Input) {
	line := command.DeletePreAndSufSpace(input.Option)
	if line == "" {
		command.PrintNotice("Invalid enter")
		GamePokerPlay(ctx, input.Data)
	} else {
		if strings.ToUpper(line) == "PASS" {
			CleanInput(input)
			ctx.pushToServer(SERVER_CODE_GAME_POKER_PLAY_PASS, "")
		} else if strings.ToUpper(line) == "EXIT" {
			CleanInput(input)
			ctx.pushToServer(SERVER_CODE_CLIENT_EXIT, "")
		} else {
			strs := strings.Split(line, " ")
			options := make([]string, 0)
			access := true
			for i := 0; i < len(strs); i++ {
				str := strs[i]
				for _, v := range []byte(str) {
					if string(v) == " " || string(v) == "\t" {
					} else {
						if !pokerLevelAliasContainer(v) {
							access = false
							break
						} else {
							options = append(options, string(v))
						}
					}
				}
			}
			if access {
				CleanInput(input)
				bytes, _ := json.Marshal(&options)
				ctx.pushToServer(SERVER_CODE_GAME_POKER_PLAY, string(bytes))
			} else {
				command.PrintNotice("Invalid enter")
				/*if ctx.LastPokers != nil {
					command.PrintNotice(ctx.LastSellClientNickname + "[" + ctx.LastSellClientType + "] playd:")
					command.PrintPokers(*ctx.LastPokers, ctx.PokerPrinterType)
				}*/
				GamePokerPlay(ctx, input.Data)
			}
		}
	}

}

func pokerLevelAliasContainer(b byte) bool {
	pokerAlias := []string{"3", "4", "5", "6", "7", "8", "9", "T", "t", "0", "J", "j", "Q", "q", "K", "k", "A", "a", "1", "2", "S", "s", "X", "x"}
	for _, v := range pokerAlias {
		if v == string(b) {
			return true
		}
	}
	return false
}
