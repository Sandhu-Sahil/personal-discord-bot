package cmd

import (
	"sandhu-sahil/bot/framework"
)

func QueueRemoveCommandIntractions(ctx *framework.Context, index int) string {
	sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
	if sess == nil {
		return ("Not in a voice channel! To make the bot join one, use `/join`.")
	}
	if !sess.Queue.Running {
		return ("Session is not running")
	}
	if index < 1 || index > sess.Queue.Length() {
		return ("Please provide a valid number, use `/queue` to see the queue")
	}

	str := sess.Queue.Remove(index - 1)
	return ("Removed from queue: `" + str + "`")
}
