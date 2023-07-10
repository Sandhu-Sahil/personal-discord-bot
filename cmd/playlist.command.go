package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sandhu-sahil/bot/framework"
	"strings"
)

func YoutubePlaylistCommandIntractions(ctx *framework.Context, url string) string {
	response := bytes.NewBufferString("")

	// error checking if session is present
	sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
	if sess == nil {
		return "No session for this server, please add me to a voice channel to start a session `/join`"
	}

	// if session exists but user not in vc
	vc := ctx.GetVoiceChannel()
	if vc == nil {
		return "Please join a voice chat to use this command"
	}
	if vc.ID != sess.ChannelId {
		return "Panic, either we are not in same vc or your vc state is empty"
	}

	// extract playlist id from url
	playlistId, err := ctx.Youtube.ExtractPlaylistId(url)
	if err != nil {
		return "Panic, error extracting playlist id: " + err.Error()
	}
	url = "https://www.youtube.com/playlist?list=" + playlistId
	output, err := ctx.Youtube.GetFromYTPlaylist(url)
	if err != nil {
		return "Panic, song extraction error: " + err.Error()
	}
	if output == "" {
		return "Panic, no songs found"
	}
	response.WriteString("Link to playlist: " + url + "\n\n")
	response.WriteString("```")
	response.WriteString("Songs added to queue:\n")

	data := strings.Split(output, "\n")
	// add songs to queue
	// start a go routine as don't want to block the main thread
	go func() {
		for index, line := range data {
			fmt.Println(line)
			if len(line) == 0 {
				ctx.Reply("Panic, no songs found")
			}
			if line == "" || line == "\n" {
				continue
			}
			var video framework.PlaylistVideo
			err := json.Unmarshal([]byte(line), &video)
			if err != nil {
				ctx.Reply(err.Error())
			}

			ctx.Youtube.Search.Id = video.Id
			types, outputTemp, err := ctx.Youtube.GetFromYT()
			if err != nil {
				ctx.Reply(err.Error())
			}
			if types == framework.ERROR_TYPE {
				ctx.Reply(*outputTemp)
			}
			if types != framework.VIDEO_TYPE {
				ctx.Reply("Panic, not a video")
			}

			mainVideo, err := ctx.Youtube.Video(*outputTemp)
			if err != nil {
				ctx.Reply("Panic, reading json: " + err.Error())
			}
			song := framework.NewSong(mainVideo.Media, mainVideo.Title, ctx.Youtube.Search.Id)
			sess.Queue.Add(*song)
			if index == 0 {
				if !sess.Queue.Running {
					go sess.Queue.Start(sess, func(msg string) {
						ctx.Reply(msg)
					})
				}
			}
		}
	}()

	response.WriteString("```")
	return response.String()
}
