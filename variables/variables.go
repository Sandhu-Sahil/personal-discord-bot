package variables

import (
	"sandhu-sahil/bot/framework"

	"github.com/bwmarrin/discordgo"
)

var (
	BotID          string
	Bot            *discordgo.Session
	Message        *discordgo.MessageCreate
	err            error
	RemoveCommands bool
	Sessions       *framework.SessionManager
	// youtube    *framework.Youtube
	PREFIX     string
	ServiceUrl string
	OwnerId    string
	// UseSharding     bool
	// ShardId         int
	// ShardCount      int
	DefaultStatus   string
	GuildID         string
	CreatedCommands []*discordgo.ApplicationCommand
)
