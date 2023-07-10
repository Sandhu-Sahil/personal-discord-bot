package cmd

import (
	"bytes"
	"fmt"
	"sandhu-sahil/bot/framework"
)

func QueueCommandIntractions(ctx *framework.Context) string {
	// retrun list of songs in queue
	sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
	if sess == nil {
		return ("Not in a voice channel! To make the bot join one, use `/join`.")
	}
	if !sess.Queue.HasNext() {
		return ("Queue is already empty")
	}
	buffer := bytes.NewBufferString("```")
	buffer.WriteString("Queue: ")

	for i, song := range sess.Queue.List() {
		if len(song.Title) > 65 {
			song.Title = song.Title[:65] + "..."
		}
		buffer.WriteString(fmt.Sprintf("\n \t%d. %s", i+1, song.Title))
	}
	buffer.WriteString("```")

	// check length of buffer
	if len(buffer.String()) > 2000 {
		// retunr that part of buffer
		return buffer.String()[:1985] + "```" + "more..."
	}

	return buffer.String()
}
