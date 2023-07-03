package cmd

import "sandhu-sahil/bot/framework"

func ClearQueueCommandIntractions(ctx *framework.Context) string {
	sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
	if sess == nil {
		return ("Not in a voice channel! To make the bot join one, use `music join`.")
	}
	if !sess.Queue.HasNext() {
		return ("Queue is already empty")
	}
	sess.Queue.Clear()
	return ("Cleared the song queue")
}
