package cmd

import (
	"sandhu-sahil/bot/framework"
	"sandhu-sahil/bot/variables"

	"github.com/bwmarrin/discordgo"
)

func YoutubeCommandIntractions(ctx *framework.Context, query string) (*[]*discordgo.MessageEmbed, string) {
	// error checking if session is present
	sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
	if sess == nil {
		return nil, "No session for this server, please add me to a voice channel to start a session `/join`"
	}

	// if session exists but user not in vc
	vc := ctx.GetVoiceChannel()
	if vc == nil {
		return nil, "Please join a voice chat to use this command"
	}
	if vc.ID != sess.ChannelId {
		return nil, "Panic, either we are not in same vc or your vc state is empty"
	}

	err := ctx.Youtube.SearchYoutube(variables.YoutubeService, query)
	if err != nil {
		return nil, "Panic, error from youtube: " + err.Error()
	}
	types, output, err := ctx.Youtube.GetFromYT()
	if err != nil {
		return nil, "Panic, song extraction error: " + err.Error()
	}
	if types != framework.VIDEO_TYPE {
		return nil, "Panic, not a video"
	}

	video, err := ctx.Youtube.Video(*output)
	if err != nil {
		return nil, "Panic, reading json: " + err.Error()
	}
	song := framework.NewSong(video.Media, video.Title, ctx.Youtube.Search.Id)
	sess.Queue.Add(*song)
	if !sess.Queue.Running {
		go sess.Queue.Start(sess, func(msg string) {
			ctx.Reply(msg)
		})
	}
	embed := ctx.CreateYoutubeEmbed()
	return embed, ""
}
