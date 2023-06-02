package handler

import (
	"log"
	"os"
	"sandhu-sahil/bot/variables"

	"github.com/bwmarrin/discordgo"
)

func StartBot() {
	var err error

	variables.Bot, err = discordgo.New("Bot " + os.Getenv("MyToken"))
	if err != nil {
		log.Fatalf("err creating bot: %v", err)
	}

	user, err := variables.Bot.User("@me")
	if err != nil {
		log.Fatal(err)
		return
	}
	variables.BotID = user.ID
}
