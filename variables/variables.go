package variables

import (
	"sandhu-sahil/bot/framework"

	"github.com/bwmarrin/discordgo"
	"google.golang.org/api/youtube/v3"
)

var (
	BotID           string
	Bot             *discordgo.Session
	Message         *discordgo.MessageCreate
	err             error
	RemoveCommands  bool
	Sessions        *framework.SessionManager
	YoutubeApiKey   string
	YoutubeService  *youtube.Service
	PREFIX          string
	ServiceUrl      string
	OwnerId         string
	DefaultStatus   string
	GuildID         string
	CreatedCommands []*discordgo.ApplicationCommand
)
