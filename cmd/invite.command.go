package cmd

import "github.com/bwmarrin/discordgo"

func InviteCommandIntractions(s *discordgo.Session, i *discordgo.InteractionCreate) string {
	// invite bot link
	// return "Invite Link: https://discord.com/api/oauth2/authorize?client_id=&permissions=8&scope=bot%20applications.commands"
	return "Invite Link: https://discord.com/api/oauth2/authorize?client_id=1023643096375898187&permissions=8&scope=bot"
}
