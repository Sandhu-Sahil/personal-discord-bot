package handler

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func RoleTrigger(s *discordgo.Session, event *discordgo.GuildCreate) {
	guildID := event.Guild.ID

	// Create the role
	roles, err := s.GuildRoles(guildID)
	if err != nil {
		log.Fatalf("Failed to create role: %v", err)
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
			log.Fatalf("Failed to create role: %v", err)
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
		log.Fatalf("Failed to update role: %v", err)
		return
	}

	// Add the role to the bot
	err = s.GuildMemberRoleAdd(guildID, s.State.User.ID, role.ID)
	if err != nil {
		log.Fatalf("Failed to add role to bot: %v", err)
		return
	}
}
