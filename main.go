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

	// delete all the files in the imports folder
	if _, err := os.Stat("./imports"); err == nil {
		err := os.RemoveAll("imports")
		if err != nil {
			log.Println("Cannot delete imports folder: " + err.Error())
		}
		// create the imports folder again with a temp file
		err = os.Mkdir("./imports", 0755)
		if err != nil {
			log.Println("Cannot create imports folder: " + err.Error())
		}
		_, err = os.Create("./imports/.temp")
		if err != nil {
			log.Println("Cannot create temp file in imports folder: " + err.Error())
		}
	}

	if variables.RemoveCommands {
		for _, cmd := range variables.CreatedCommands {
			err := variables.Bot.ApplicationCommandDelete(variables.BotID, "", cmd.ID) // if guildID is empty "", it will create global commands
			if err != nil {
				log.Fatalf("Cannot delete %q command: %v", cmd.Name, err)
			}
			fmt.Printf("Command %q has been deleted\n", cmd.Name)
		}
	}

	log.Println("Bot is now offline. Goodbye!")
	fmt.Println()
}
