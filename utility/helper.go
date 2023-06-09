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

	handler.StartYoutubeClient()

	variables.Sessions = framework.NewSessionManager()

	// if variables.UseSharding {
	// 	variables.Bot.ShardID = variables.ShardId
	// 	variables.Bot.ShardCount = variables.ShardCount
	// }

	// raw message handler
	variables.Bot.AddHandler(handler.MessageHandler)

	// variables.Bot.Identify.Intents = discordgo.IntentsGuildMessages
	// variables.Bot.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildVoiceStates | discordgo.IntentsGuildMembers | discordgo.IntentsGuildPresences | discordgo.IntentsGuildMessages
	variables.Bot.ShardID = 0
	variables.Bot.ShardCount = 1
	// this same as setting up Intents in the bot

	// command handler
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
	variables.CreatedCommands, err = variables.Bot.ApplicationCommandBulkOverwrite(variables.BotID, "", handler.Commands) // if guildID is empty "", it will create global commands
	if err != nil {
		log.Fatalf("Cannot register commands: %v", err)
	}

	err = variables.Bot.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
		return
	}

	// create a role for bot in new server
	variables.Bot.AddHandler(handler.RoleTrigger)

}
