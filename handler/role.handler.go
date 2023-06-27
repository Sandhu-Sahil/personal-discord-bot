package handler

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func RoleTrigger(s *discordgo.Session, event *discordgo.GuildCreate) {
	// check if the bot has the correct permissions
	permissions, err := s.UserChannelPermissions(s.State.User.ID, event.Guild.SystemChannelID)
	if err != nil {
		permissions, err = s.UserChannelPermissions(s.State.User.ID, event.Guild.Channels[0].ID)
		if err != nil {
			fmt.Printf("Failed to get channel permissions in %s: %v", event.Guild.Name, err)
			return
		}
	}

	// if the bot doesn't have administrator permissions create a role without administrator permissions
	if permissions&discordgo.PermissionAdministrator == 0 {
		// instead of returning send message to the server system channel
		_, err = s.ChannelMessageSend(event.Guild.SystemChannelID, "I don't have administrator permissions in this server. Please give me administrator permissions so that I can create a role for myself.")
		if err != nil {
			fmt.Println("Failed to send message to the system channel")
			// send message to the first channel of the server
			_, err = s.ChannelMessageSend(event.Guild.Channels[0].ID, "I don't have administrator permissions in this server. Please give me administrator permissions so that I can create a role for myself.")
			if err != nil {
				fmt.Printf("Failed to send message to the first channel of the server: %v", err)
				return
			}
		}
	} else {
		guildID := event.Guild.ID

		// Create the role
		roles, err := s.GuildRoles(guildID)
		if err != nil {
			fmt.Printf("Failed to create role: %v", err)
			return
		}

		// Find the role
		var role *discordgo.Role
		for _, r := range roles {
			if r.Name == "Sandhu Saab" {
				role = r
				break
			}
		}

		// Create a new role if it doesn't exist
		if role == nil {
			role, err = s.GuildRoleCreate(guildID, &discordgo.RoleParams{})
			if err != nil {
				fmt.Printf("Failed to create role in %s: %v", event.Guild.Name, err)
				return
			}
			// Customize the role
			role.Name = "Sandhu Saab"
			role.Permissions = discordgo.PermissionAdministrator // Set desired permissions
			role.Color = 0x89CFF0                                // Set desired color
			role.Hoist = true                                    // Show users with this role separately in the sidebar
			role.Mentionable = true                              // Allow anyone to mention this role
		}

		// Update the role with customized settings
		_, err = s.GuildRoleEdit(guildID, role.ID, &discordgo.RoleParams{
			Name:        role.Name,
			Permissions: &role.Permissions,
			Color:       &role.Color,
			Hoist:       &role.Hoist,
			Mentionable: &role.Mentionable,
		})
		if err != nil {
			fmt.Printf("Failed to update role: %v", err)
			return
		}

		// Add the role to the bot
		err = s.GuildMemberRoleAdd(guildID, s.State.User.ID, role.ID)
		if err != nil {
			fmt.Printf("Failed to add role to bot: %v", err)
			return
		}
	}
}
