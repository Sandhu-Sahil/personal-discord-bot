package handler

import (
	"github.com/bwmarrin/discordgo"
)

func CommandHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	// user := message.Author
	// if user.ID == variables.BotID || user.Bot {
	// 	return
	// }
	// content := message.Content
	// if len(content) <= len(variables.PREFIX) {
	// 	return
	// }
	// if content[:len(variables.PREFIX)] != variables.PREFIX {
	// 	return
	// }
	// content = content[len(variables.PREFIX):]
	// if len(content) < 1 {
	// 	return
	// }
	// args := strings.Fields(content)
	// name := strings.ToLower(args[0])
	// command, found := variables.CmdHandler.Get(name)
	// if !found {
	// 	fmt.Println("Command not found,", name)
	// 	return
	// }
	// channel, err := session.Channel(message.ChannelID)
	// if err != nil {
	// 	fmt.Println("Error getting channel,", err)
	// 	return
	// }
	// guild, err := session.State.Guild(channel.GuildID)
	// if err != nil {
	// 	fmt.Println("Error getting guild,", err)
	// 	return
	// }
	// ctx := framework.NewContext(session, guild, channel, user) // Sessions, youtube

	// ctx.Args = args[1:]
	// c := *command
	// c(*ctx)
}

func RegisterCommands() {
	// ??? means I haven't dug in
	// TODO: Consistant order?
	// variables.CmdHandler.Register("help", cmd.HelpCommand, "Gives you this help message!")
	// variables.CmdHandler.Register("admin", cmd.AdminCommand, "???")
	// variables.CmdHandler.Register("join", cmd.JoinCommand, "Join a voice channel !join attic")
	// variables.CmdHandler.Register("leave", cmd.LeaveCommand, "Leaves current voice channel")
	// variables.CmdHandler.Register("play", cmd.PlayCommand, "Plays whats in the queue")
	// variables.CmdHandler.Register("stop", cmd.StopCommand, "Stops the music")
	// variables.CmdHandler.Register("info", cmd.InfoCommand, "???")
	// variables.CmdHandler.Register("add", cmd.AddCommand, "Add a song to the queue !add <youtube-link>")
	// variables.CmdHandler.Register("skip", cmd.SkipCommand, "Skip")
	// variables.CmdHandler.Register("queue", cmd.QueueCommand, "Print queue???")
	// variables.CmdHandler.Register("eval", cmd.EvalCommand, "???")
	// variables.CmdHandler.Register("debug", cmd.DebugCommand, "???")
	// variables.CmdHandler.Register("clear", cmd.ClearCommand, "empty queue???")
	// variables.CmdHandler.Register("current", cmd.CurrentCommand, "Name current song???")
	// variables.CmdHandler.Register("youtube", cmd.YoutubeCommand, "???")
	// variables.CmdHandler.Register("shuffle", cmd.ShuffleCommand, "Shuffle queue???")
	// variables.CmdHandler.Register("pausequeue", cmd.PauseCommand, "Pause song in place???")
	// variables.CmdHandler.Register("pick", cmd.PickCommand, "???")
}
