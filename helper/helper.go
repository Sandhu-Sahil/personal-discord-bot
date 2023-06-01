package helper

import (
	"sandhu-sahil/bot/handler"
	"sandhu-sahil/bot/variables"

	"github.com/bwmarrin/discordgo"
)

func HandleBot() {
	variables.Bot.AddHandler(handler.MessageHandler)
	variables.Bot.Identify.Intents = discordgo.IntentsGuildMessages
	err := variables.Bot.Open()
	if err != nil {
		return
	}
}
