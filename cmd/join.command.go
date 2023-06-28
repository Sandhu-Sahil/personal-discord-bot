package cmd

import "sandhu-sahil/bot/framework"

func JoinCommandIntractions(ctx *framework.Context) string {
	if ctx.Sessions.GetByGuild(ctx.Guild.ID) != nil {
		return "Already connected! Use `/leave` for the bot to disconnect."
	}
	vc := ctx.GetVoiceChannel()
	if vc == nil {
		return "You must be in a voice channel to use the bot!"
	}
	sess, err := ctx.Sessions.Join(ctx.Discord, ctx.Guild.ID, vc.ID, framework.JoinProperties{
		Muted:    false,
		Deafened: true,
	})
	if err != nil {
		return "An error occured at the time of joining! Please try again later "
	}
	return "Joined <#" + sess.ChannelId + ">!"
}
