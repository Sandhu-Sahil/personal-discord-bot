package main

import (
	"fmt"
	"log"
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

	defer variables.Bot.Close()

	log.Println("Bot is now up and running. Press CTRL-C to exit :)")
	//line break
	fmt.Println()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	fmt.Println()
	log.Println("Gracefully shutting down, Please wait...")

	if variables.RemoveCommands {
		for _, cmd := range variables.CreatedCommands {
			err := variables.Bot.ApplicationCommandDelete(variables.BotID, "", cmd.ID) // if guildID is empty "", it will create global commands
			if err != nil {
				log.Fatalf("Cannot delete %q command: %v", cmd.Name, err)
			}
		}
	}

	log.Println("Bot is now offline. Goodbye!")
	fmt.Println()
}
