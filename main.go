package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}
}

func main() {
	Bot, err = NewBot()
	if err != nil {
		log.Fatalf("err creating bot: %v", err)
	}

	user, err := Bot.User("@me")
	if err != nil {
		log.Fatal(err)
		return
	}
	BotID = user.ID

	HandleBot()

	fmt.Println("Bot is now running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	Bot.Close()
}
