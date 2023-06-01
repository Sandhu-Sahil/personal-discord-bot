package variables

import "github.com/bwmarrin/discordgo"

var (
	BotID   string
	Bot     *discordgo.Session
	Message *discordgo.MessageCreate
	err     error
)
