package cmd

import (
	"fmt"
	"sandhu-sahil/bot/framework"
)

func LoopQueueCommandIntractions(ctx *framework.Context) string {
	sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
	if sess == nil {
		return ("Not in a voice channel! To make the bot join one, use `/join`.")
	}
	if !sess.Queue.Running {
		return ("Session is not running")
	}
	// add current song to queue and then toogle loop queue
	if !sess.Queue.LoopQueue {
		sess.Queue.Add(*sess.Queue.Current())
	}
	sess.Queue.ToogleLoopQueue()
	return ("Toggled loop queue to " + fmt.Sprint(sess.Queue.LoopQueue) + ".")
}
