package main

import (
	"fmt"
	"os"
	"os/signal"
	"sandhu-sahil/bot/utility"
	"sandhu-sahil/bot/variables"
	"syscall"
)

func main() {
	utility.HandleBot()

	// can use this also
	// <-make(chan struct{})

	fmt.Println("Bot is now running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	variables.Bot.Close()
}
