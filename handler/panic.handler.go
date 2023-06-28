package handler

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func BotPanicHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if r := recover(); r != nil {
		// Log the panic error
		log.Println("Panic occurred:", r)

		str := fmt.Sprintln("Panic occurred: " + fmt.Sprintln(r))
		_, err := s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &str,
		})

		if err != nil {
			log.Println("Recoverd from panic but unable to update intraction :]")
		} else {
			log.Println("Fully recoved from panic successfully :)")
		}
	}
}
