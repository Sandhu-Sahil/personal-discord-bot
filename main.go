package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sandhu-sahil/bot/helper"
	"sandhu-sahil/bot/variables"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

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

func main() {
	helper.HandleBot()

	fmt.Println("Bot is now running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	variables.Bot.Close()
}
