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
	case strings.Contains(strings.ToUpper(m.Content), "LTV"):
		{
			_, err := s.ChannelMessageSend(m.ChannelID, "tu LTV panchooda")
			if err != nil {
				fmt.Print(err)
			}
		}
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
		// case strings.Contains(strings.ToUpper(m.Content), "HI"):
		// 	{
		// 		_, err := s.ChannelMessageSend(m.ChannelID, "Hi sir, how can I help you!")
		// 		if err != nil {
		// 			fmt.Print(err)
		// 		}
		// 	}
		// instead of above case we can use below case
		// instead of checking inbetween words we can check the whole word
		// that is why we are using space before and after the word
		// and also check if that only same word is present in the message
	case strings.Contains(strings.ToUpper(m.Content), "HI"):
		{
			switch {

			// if the message is only hi
			case strings.ToUpper(m.Content) == "HI":
				{
					_, err := s.ChannelMessageSend(m.ChannelID, "Hi sir, how can I help you!")
					if err != nil {
						fmt.Print(err)
					}
				}
			// if hi is the first word in the message
			case strings.HasPrefix(strings.ToUpper(m.Content), "HI "):
				{
					// check if hi is really the first word or not
					// if the message is "hi there" not "yo meetchi there"
					if strings.Split(strings.ToUpper(m.Content), " ")[0] == "HI" {
						_, err := s.ChannelMessageSend(m.ChannelID, "Hi sir, how can I help you!")
						if err != nil {
							fmt.Print(err)
						}
					}
				}
			// if hi is the last word in the message
			case strings.HasSuffix(strings.ToUpper(m.Content), " HI"):
				{
					// check if hi is really the last word or not
					// if the message is "there hi" not "there hit meetchi"
					if strings.Split(strings.ToUpper(m.Content), " ")[len(strings.Split(strings.ToUpper(m.Content), " "))-1] == "HI" {
						_, err := s.ChannelMessageSend(m.ChannelID, "Hi sir, how can I help you!")
						if err != nil {
							fmt.Print(err)
						}
					}
				}
			// if the message is hi in between some words
			case strings.Contains(strings.ToUpper(m.Content), " HI "):
				{
					_, err := s.ChannelMessageSend(m.ChannelID, "Hi sir, how can I help you!")
					if err != nil {
						fmt.Print(err)
					}
				}
			}

		}

	// instead of what i have done above we can use below case
	// strings.split will split the message into words
	// and then we can check if the word is present in the message or not
	// but this is a bit slower than the above case
	// so we will use the above case
	case strings.Contains(strings.ToUpper(m.Content), "BYE"):
		{
			_, err := s.ChannelMessageSend(m.ChannelID, "Bye sir, have a nice time ahead!")
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
	if strings.Contains(m.ContentWithMentionsReplaced(), `@deathonn`) { // personal
		_, err := s.ChannelMessageSend(m.ChannelID, "peo nu kato vaja marre janda")
		if err != nil {
			fmt.Print(err)
		}
	}
}
