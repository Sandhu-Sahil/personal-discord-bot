package handler

import "github.com/bwmarrin/discordgo"

var Commands = []*discordgo.ApplicationCommand{
	{
		Name:        "help",
		Description: "Showcase of help command",
		Type:        discordgo.ChatApplicationCommand,
	},
	{
		Name:        "join",
		Description: "Join the voice channel",
		Type:        discordgo.ChatApplicationCommand,
	},
	{
		Name:        "leave",
		Description: "Leave the voice channel",
		Type:        discordgo.ChatApplicationCommand,
	},
	{
		Name:        "support",
		Description: "Showcase of support server link",
		Type:        discordgo.ChatApplicationCommand,
	},
	{
		Name:        "single-autocomplete",
		Description: "Showcase of single autocomplete option",
		Type:        discordgo.ChatApplicationCommand,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:         "autocomplete-option",
				Description:  "Autocomplete option",
				Type:         discordgo.ApplicationCommandOptionString,
				Required:     true,
				Autocomplete: true,
			},
		},
	},
	{
		Name:        "multi-autocomplete",
		Description: "Showcase of multiple autocomplete option",
		Type:        discordgo.ChatApplicationCommand,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:         "autocomplete-option-1",
				Description:  "Autocomplete option 1",
				Type:         discordgo.ApplicationCommandOptionString,
				Required:     true,
				Autocomplete: true,
			},
			{
				Name:         "autocomplete-option-2",
				Description:  "Autocomplete option 2",
				Type:         discordgo.ApplicationCommandOptionString,
				Required:     true,
				Autocomplete: true,
			},
		},
	},
}
