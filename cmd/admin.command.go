package cmd

import (
	"sandhu-sahil/bot/variables"

	"github.com/bwmarrin/discordgo"
)

func AdminCommandIntractions(s *discordgo.Session, i *discordgo.InteractionCreate) string {
	// check if it requester is owner of bot
	if i.Member.User.ID != variables.OwnerId {
		return "You are not authorized to use this command"
	}
	return "Will create Admin Intractions when needed"
}
