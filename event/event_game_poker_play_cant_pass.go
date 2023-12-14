package event

import "fish/client/command"

func GamePokerPlayCantPass(ctx *EContext, data string) {
	command.PrintNotice("You played the previous card, so you can't pass.")
	ctx.pushToServer(SERVER_CODE_GAME_POKER_PLAY_REDIRECT, "")
}
