package utility

import (
	"fmt"
	"sandhu-sahil/bot/framework"
	"sandhu-sahil/bot/handler"
	"sandhu-sahil/bot/variables"

	"github.com/bwmarrin/discordgo"
)

func HandleBot() {
	handler.LoadEnv()

	handler.StartBot()

	variables.CmdHandler = framework.NewCommandHandler()
	handler.RegisterCommands()

	// if variables.UseSharding {
	// 	variables.Bot.ShardID = variables.ShardId
	// 	variables.Bot.ShardCount = variables.ShardCount
	// }

	variables.Bot.AddHandler(handler.CommandHandler)
	variables.Bot.AddHandler(handler.MessageHandler)
	variables.Bot.Identify.Intents = discordgo.IntentsGuildMessages
	variables.Bot.AddHandler(func(discord *discordgo.Session, ready *discordgo.Ready) {
		discord.UpdateWatchStatus(0, variables.DefaultStatus)
		guilds := discord.State.Guilds
		fmt.Println("Ready with", len(guilds), "guilds.")
	})
	err := variables.Bot.Open()
	if err != nil {
		return
	}
}
