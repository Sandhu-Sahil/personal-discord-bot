package framework

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func ExtractDataCreateContext(s *discordgo.Session, i *discordgo.InteractionCreate, session *SessionManager, keyYoutube string) (ctx *Context, err error) {
	channel, err := s.Channel(i.ChannelID)
	if err != nil {
		fmt.Println("Error getting channel,", err)
		return nil, err
	}
	guild, err := s.State.Guild(i.GuildID)
	if err != nil {
		fmt.Println("Error getting guild,", err)
		return nil, err
	}
	user, err := s.User(i.Member.User.ID)
	if err != nil {
		fmt.Println("Error getting user,", err)
		return nil, err
	}
	youtube := NewYoutube(keyYoutube)
	ctx = NewContext(s, guild, channel, user, session, youtube)
	return ctx, nil
}
