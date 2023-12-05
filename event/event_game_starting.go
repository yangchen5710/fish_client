package event

import (
	"encoding/json"
	"fish/client/command"
)

func GameStarting(ctx *EContext, data string) {
	command.PrintNotice("Game starting !!")

	pokers := make([]Poker, 0)
	_ = json.Unmarshal([]byte(data), &pokers)

	command.PrintNotice("")
	command.PrintNotice("Your pokers are")
	command.PrintPokers(pokers, ctx.PokerPrinterType)

	GameLandlordElect(ctx, data)
}
