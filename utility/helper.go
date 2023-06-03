package utility

import (
	"fmt"
	"log"
	"sandhu-sahil/bot/framework"
	"sandhu-sahil/bot/handler"
	"sandhu-sahil/bot/variables"

	"github.com/bwmarrin/discordgo"
)

func HandleBot() {
	var err error

	handler.LoadEnv()

	handler.StartBot()

	variables.CmdHandler = framework.NewCommandHandler()
	handler.RegisterCommands()

	// if variables.UseSharding {
	// 	variables.Bot.ShardID = variables.ShardId
	// 	variables.Bot.ShardCount = variables.ShardCount
	// }

	// raw message handler
	variables.Bot.AddHandler(handler.MessageHandler)
	variables.Bot.Identify.Intents = discordgo.IntentsGuildMessages

	// command handler
	variables.Bot.AddHandler(handler.CommandHandler)
	variables.Bot.AddHandler(func(discord *discordgo.Session, ready *discordgo.Ready) {
		discord.UpdateWatchStatus(0, variables.DefaultStatus)
		guilds := discord.State.Guilds
		fmt.Println("Ready with", len(guilds), "guilds.")
	})

	// register the interactions
	variables.Bot.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := handler.IntractionHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
	variables.CreatedCommands, err = variables.Bot.ApplicationCommandBulkOverwrite(variables.BotID, variables.GuildID, handler.Commands)
	if err != nil {
		log.Fatalf("Cannot register commands: %v", err)
	}

	err = variables.Bot.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
		return
	}
}
