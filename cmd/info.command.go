package cmd

import (
	"bytes"
	"fmt"
	"runtime"
	"sandhu-sahil/bot/framework"
	"sandhu-sahil/bot/variables"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/dustin/go-humanize"
)

var startTime = time.Now()
var userString *string

func getDurationString(duration time.Duration) string {
	return fmt.Sprintf(
		"%0.2d:%02d:%02d",
		int(duration.Hours()),
		int(duration.Minutes())%60,
		int(duration.Seconds())%60,
	)
}

func InfoCommandIntractions(ctx *framework.Context) string {
	if userString == nil {
		usr, err := ctx.Discord.User(variables.OwnerId)
		if err != nil {
			fmt.Println("error getting user ", variables.OwnerId, err)
			return "error getting user " + variables.OwnerId + err.Error()
		}
		str := usr.Username
		userString = &str
	}
	stats := runtime.MemStats{}
	runtime.ReadMemStats(&stats)
	buffer := bytes.NewBufferString("```")

	buffer.WriteString("owner username: " + *userString)
	buffer.WriteString("\ngo version: " + runtime.Version())
	buffer.WriteString("\ndiscordgo version: " + discordgo.VERSION)
	buffer.WriteString("\nuptime: " + getDurationString(time.Now().Sub(startTime)))
	buffer.WriteString(fmt.Sprintf("\nmemory used: %s / %s (%s garbage collected)", humanize.Bytes(stats.Alloc),
		humanize.Bytes(stats.Sys), humanize.Bytes(stats.TotalAlloc)))
	buffer.WriteString("\nconcurrent tasks: " + strconv.Itoa(runtime.NumGoroutine()))
	buffer.WriteString("```")
	return buffer.String()
}
