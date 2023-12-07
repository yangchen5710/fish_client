package event

import "fish/client/command"

func GameLandlordCycle(ctx *EContext, data string) {
	command.PrintNotice("No player takes the landlord, so redealing cards.")
}
