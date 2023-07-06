package cmd

import "sandhu-sahil/bot/framework"

func CurrentSongCommandIntractions(ctx *framework.Context) string {
	sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
	if sess == nil {
		return ("Not in a voice channel! To make the bot join one, use `/join`.")
	}
	current := sess.Queue.Current()
	if current == nil {
		return ("The queue is empty! Add with `/youtube`.")
	}
	return ("Currently playing `" + current.Title + "`.")
}
