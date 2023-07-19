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
		Name:        "youtube",
		Description: "Play youtube video",
		Type:        discordgo.ChatApplicationCommand,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "search",
				Description: "Youtube search query",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
		},
	},
	{
		Name:        "skip",
		Description: "Skip the current song",
		Type:        discordgo.ChatApplicationCommand,
	},
	{
		Name:        "clear-queue",
		Description: "Clear the songs queue",
		Type:        discordgo.ChatApplicationCommand,
	},
	{
		Name:        "currently-playing",
		Description: "Showcase of current playing song",
		Type:        discordgo.ChatApplicationCommand,
	},
	{
		Name:        "queue",
		Description: "Showcase of queue",
		Type:        discordgo.ChatApplicationCommand,
	},
	{
		Name:        "pause",
		Description: "Pause the current song",
		Type:        discordgo.ChatApplicationCommand,
	},
	{
		Name:        "resume",
		Description: "Resume the current song",
		Type:        discordgo.ChatApplicationCommand,
	},
	{
		Name:        "repeat",
		Description: "Toggle repeat for current song",
		Type:        discordgo.ChatApplicationCommand,
	},
	{
		Name:        "loop-queue",
		Description: "Toggle loop for current queue",
		Type:        discordgo.ChatApplicationCommand,
	},
	{
		Name:        "replay",
		Description: "Replay the current song",
		Type:        discordgo.ChatApplicationCommand,
	},
	{
		Name:        "queue-remove",
		Description: "Remove a song from queue",
		Type:        discordgo.ChatApplicationCommand,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "index",
				Description: "Index of song to remove",
				Type:        discordgo.ApplicationCommandOptionInteger,
				Required:    true,
			},
		},
	},
	{
		Name:        "youtube-playlist-search",
		Description: "Play youtube playlist",
		Type:        discordgo.ChatApplicationCommand,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "youtube-url",
				Description: "Youtube playlist url",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
		},
	},
	{
		Name:        "admin",
		Description: "Admin only commands",
		Type:        discordgo.ChatApplicationCommand,
	},
	{
		Name:        "info",
		Description: "Showcase of info command",
		Type:        discordgo.ChatApplicationCommand,
	},
	{
		Name:        "invite",
		Description: "Invite the bot to your server",
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
