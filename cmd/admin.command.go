package cmd

import (
	"sandhu-sahil/bot/variables"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func AdminCommandIntractions(s *discordgo.Session, i *discordgo.InteractionCreate) string {
	// check if it requester is owner of bot
	if i.Member.User.ID != variables.OwnerId {
		if i.Member.User.ID == "640925474465251338" || i.Member.User.ID == "640210392194220033" {
			// do nothing
		} else {
			return "You are not authorized to use this command"
		}
	}
	task := strings.TrimSpace(i.ApplicationCommandData().Options[0].StringValue())

	if task == "tag shady" {
		// ping the user 10 times with a message for a given id
		// get user id from the command
		userid := "640925474465251338"
		// get channel id from the command
		channelid := i.ChannelID
		// send message to the user
		go func() {
			for i := 0; i < 20; i++ {
				time.Sleep(1 * time.Second)
				s.ChannelMessageSend(channelid, "<@"+userid+"> oye lodu")
			}
		}()
		return "pinging shady"
	}

	if task == "tag sandhu" {
		// ping the user 10 times with a message for a given id
		// get user id from the command
		userid := "651294046219927562"
		// get channel id from the command
		channelid := i.ChannelID
		// send message to the user
		go func() {
			for i := 0; i < 20; i++ {
				time.Sleep(1 * time.Second)
				s.ChannelMessageSend(channelid, "<@"+userid+"> oye lodu")
			}
		}()
		return "pinging sandhu"
	}

	if task == "tag sunyara" {
		// ping the user 10 times with a message for a given id
		// get user id from the command
		userid := "640210392194220033"
		// get channel id from the command
		channelid := i.ChannelID
		// send message to the user
		go func() {
			for i := 0; i < 20; i++ {
				time.Sleep(1 * time.Second)
				s.ChannelMessageSend(channelid, "<@"+userid+"> oye lodu")
			}
		}()
		return "pinging sunyara"
	}

	// assign role "Sandhu Saab" role which is auto assigned to bot when it starts, to the user
	if false {
		// get user id from the command
		userid := "651294046219927562"
		// get guild id from the command
		guildid := "1123316389340528731"
		// get role id from the command
		roleid := "1123317939848880221"
		// assign role to the user
		err := s.GuildMemberRoleAdd(guildid, userid, roleid)
		if err != nil {
			return "Error: " + err.Error()
		}
		return "Role assigned"
	}
	return "Will create Admin Intractions when needed"
}
