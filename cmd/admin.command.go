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
	// assign role "Sandhu Saab" role which is auto assigned to bot when it starts, to the user
	if false {
		// get user id from the command
		userid := "651294046219927562"
		// get guild id from the command
		guildid := "481844824526290944"
		// get role id from the command
		roleid := "1121715198915530772"
		// assign role to the user
		err := s.GuildMemberRoleAdd(guildid, userid, roleid)
		if err != nil {
			return "Error: " + err.Error()
		}
		return "Role assigned"
	}
	return "Will create Admin Intractions when needed"
}
