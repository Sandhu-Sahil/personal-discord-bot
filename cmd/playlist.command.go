package cmd

import (
	"sandhu-sahil/bot/framework"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func YoutubePlaylistCommandIntractions(ctx *framework.Context, url string) (*[]*discordgo.MessageEmbed, string) {
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

	// extract playlist id from url
	playlistId, err := ctx.Youtube.ExtractPlaylistId(url)
	if err != nil {
		return nil, "Panic, error extracting playlist id: " + err.Error()
	}
	url = "https://www.youtube.com/playlist?list=" + playlistId
	output, err := ctx.Youtube.GetFromYTPlaylist(url)
	if err != nil {
		return nil, "Panic, song extraction error: " + err.Error()
	}
	if output == "" {
		return nil, "Panic, no playlist found"
	}

	data := strings.Split(output, "\n")
	// add songs to queue
	// start a go routine as don't want to block the main thread
	go sess.Queue.StartPlaylist(data, ctx, sess, func(msg string) {
		ctx.Reply(msg)
	})

	embed := ctx.CreateYoutubePlaylistEmbed(data[0], url)

	return embed, ""
}
