package handler

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var Commands = []*discordgo.ApplicationCommand{
	{
		Name:        "help",
		Description: "Showcase of help command",
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

var IntractionHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"help": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf(
						"Help command response",
					),
				},
			})
			if err != nil {
				panic(err)
			}
		}
	},
	"single-autocomplete": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			data := i.ApplicationCommandData()
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf(
						"You picked %q autocompletion",
						// Autocompleted options do not affect usual flow of handling application command. They are ordinary options at this stage
						data.Options[0].StringValue(),
					),
				},
			})
			if err != nil {
				panic(err)
			}
		// Autocomplete options introduce a new interaction type (8) for returning custom autocomplete results.
		case discordgo.InteractionApplicationCommandAutocomplete:
			data := i.ApplicationCommandData()
			choices := []*discordgo.ApplicationCommandOptionChoice{
				{
					Name:  "Autocomplete",
					Value: "autocomplete",
				},
				{
					Name:  "Autocomplete is best!",
					Value: "autocomplete_is_best",
				},
				{
					Name:  "Choice 3",
					Value: "choice3",
				},
				{
					Name:  "Choice 4",
					Value: "choice4",
				},
				{
					Name:  "Choice 5",
					Value: "choice5",
				},
				// And so on, up to 25 choices
			}

			if data.Options[0].StringValue() != "" {
				choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
					Name:  data.Options[0].StringValue(), // To get user input you just get value of the autocomplete option.
					Value: data.Options[0].StringValue(),
				})
			}

			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionApplicationCommandAutocompleteResult,
				Data: &discordgo.InteractionResponseData{
					Choices: choices, // This is basically the whole purpose of autocomplete interaction - return custom options to the user.
				},
			})
			if err != nil {
				panic(err)
			}
		}
	},
	"multi-autocomplete": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			data := i.ApplicationCommandData()
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf(
						"Option 1: %s\nOption 2: %s",
						data.Options[0].StringValue(),
						data.Options[1].StringValue(),
					),
				},
			})
			if err != nil {
				panic(err)
			}
		case discordgo.InteractionApplicationCommandAutocomplete:
			data := i.ApplicationCommandData()
			var choices []*discordgo.ApplicationCommandOptionChoice
			switch {
			// In this case there are multiple autocomplete options. The Focused field shows which option user is focused on.
			case data.Options[0].Focused:
				choices = []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "Autocomplete 4 first option",
						Value: "autocomplete_default",
					},
					{
						Name:  "Choice 3",
						Value: "choice3",
					},
					{
						Name:  "Choice 4",
						Value: "choice4",
					},
					{
						Name:  "Choice 5",
						Value: "choice5",
					},
				}
				if data.Options[0].StringValue() != "" {
					choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
						Name:  data.Options[0].StringValue(),
						Value: "choice_custom",
					})
				}

			case data.Options[1].Focused:
				choices = []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "Autocomplete 4 second option",
						Value: "autocomplete_1_default",
					},
					{
						Name:  "Choice 3.1",
						Value: "choice3_1",
					},
					{
						Name:  "Choice 4.1",
						Value: "choice4_1",
					},
					{
						Name:  "Choice 5.1",
						Value: "choice5_1",
					},
				}
				if data.Options[1].StringValue() != "" {
					choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
						Name:  data.Options[1].StringValue(),
						Value: "choice_custom_2",
					})
				}
			}

			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionApplicationCommandAutocompleteResult,
				Data: &discordgo.InteractionResponseData{
					Choices: choices,
				},
			})
			if err != nil {
				panic(err)
			}
		}
	},
}
