package cmd

import "sandhu-sahil/bot/framework"

func ResumeCommandIntractions(ctx *framework.Context) string {
	sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
	if sess == nil {
		return ("Not in a voice channel! To make the bot join one, use `/join`.")
	}
	if !sess.Queue.Running {
		return ("Session is not running!")
	}
	sess.Resume()
	return ("Resumed")
}
