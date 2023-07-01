package cmd

import "sandhu-sahil/bot/framework"

func SkipCommandIntractions(ctx *framework.Context) string {
	sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
	if sess == nil {
		return "Not in a voice channel! To make the bot join one, use `music join`."
	}
	sess.Stop()
	return ("Skipped song!")
}
