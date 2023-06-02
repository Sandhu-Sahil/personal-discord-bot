package handler

import (
	"fmt"
	"sandhu-sahil/bot/variables"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == variables.BotID {
		return
	}

	// message contain some text and the bot mention
	switch {
	case strings.Contains(strings.ToUpper(m.Content), "BOT"):
		{
			_, err := s.ChannelMessageSend(m.ChannelID, "Hi there!")
			if err != nil {
				fmt.Print(err)
			}
		}
	case strings.Contains(strings.ToUpper(m.Content), "HELLO"):
		{
			_, err := s.ChannelMessageSend(m.ChannelID, "Hello sir, how can I help you!")
			if err != nil {
				fmt.Print(err)
			}
		}
	case strings.Contains(strings.ToUpper(m.Content), "HI"):
		{
			_, err := s.ChannelMessageSend(m.ChannelID, "Hi sir, how can I help you!")
			if err != nil {
				fmt.Print(err)
			}
		}
	case strings.Contains(strings.ToUpper(m.Content), "LTV"):
		{
			_, err := s.ChannelMessageSend(m.ChannelID, "tu LTV panchooda")
			if err != nil {
				fmt.Print(err)
			}
		}

	}

	// if message contain personal mentions
	if strings.Contains(m.ContentWithMentionsReplaced(), `<@&998929334645563502>`) { //@fudi deya
		_, err := s.ChannelMessageSend(m.ChannelID, "kithe marre aa kuriyave sare")
		if err != nil {
			fmt.Print(err)
		}
	}
	if strings.Contains(m.ContentWithMentionsReplaced(), `<@&765554608172957696>`) { //@kings
		_, err := s.ChannelMessageSend(m.ChannelID, "Acha, gusse sare king bnate aa")
		if err != nil {
			fmt.Print(err)
		}
	}
	if strings.Contains(m.ContentWithMentionsReplaced(), `<@&789483612072181771>`) { //@ace
		_, err := s.ChannelMessageSend(m.ChannelID, "aces te oda he FUDU aa")
		if err != nil {
			fmt.Print(err)
		}
	}
	if strings.Contains(m.ContentWithMentionsReplaced(), `<@&1090312377876107324>`) { //@shady sandhu sunyara squad
		_, err := s.ChannelMessageSend(m.ChannelID, "serious kaamm aa jalde kathe hovo")
		if err != nil {
			fmt.Print(err)
		}
	}
	if strings.Contains(m.ContentWithMentionsReplaced(), `@Sahil Sandhu`) { // personal
		_, err := s.ChannelMessageSend(m.ChannelID, "peo nu kato vaja marre janda")
		if err != nil {
			fmt.Print(err)
		}
	}
}
