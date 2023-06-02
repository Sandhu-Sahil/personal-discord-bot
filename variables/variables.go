package variables

import (
	"sandhu-sahil/bot/framework"

	"github.com/bwmarrin/discordgo"
)

var (
	BotID   string
	Bot     *discordgo.Session
	Message *discordgo.MessageCreate
	err     error

	// Sessions   *framework.SessionManager
	// youtube    *framework.Youtube
	CmdHandler *framework.CommandHandler
	PREFIX     string

	ServiceUrl    string
	OwnerId       string
	UseSharding   bool
	ShardId       int
	ShardCount    int
	DefaultStatus string
)
