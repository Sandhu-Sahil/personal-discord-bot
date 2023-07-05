package cmd

import "sandhu-sahil/bot/framework"

func LeaveCommandIntractions(ctx *framework.Context) string {
	sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
	if sess == nil {
		return "Not in a voice channel! To make the bot join one, use `/join`."
	}
	sess.Stop()
	ctx.Sessions.Leave(ctx.Discord, *sess)
	return "Left <#" + sess.ChannelId + ">!"
}
