package cmd

import (
	"sandhu-sahil/bot/variables"

	"github.com/bwmarrin/discordgo"
)

func SupportCommandIntractions(s *discordgo.Session, i *discordgo.InteractionCreate) string {
	// get guild
	guild, err := s.Guild(variables.GuildID)
	if err != nil {
		return "Error: " + err.Error()
	}
	// create server invite link
	invite, err := s.ChannelInviteCreate(guild.RulesChannelID, discordgo.Invite{
		MaxAge:  86400,
		MaxUses: 10,
	})
	if err != nil {
		return "Error: " + err.Error()
	}
	// send invite link
	return "Invite Link: https://discord.gg/" + invite.Code
}
