package framework

import (
	"encoding/json"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type Embed struct {
	PlayListCount     int    `json:"playlist_count"`
	PlaylistTitle     string `json:"playlist_title"`
	PlaylistUploader  string `json:"playlist_uploader"`
	PlaylistThumbnail []struct {
		Url string `json:"url"`
	} `json:"thumbnails"`
}

type Context struct {
	Discord      *discordgo.Session
	Guild        *discordgo.Guild
	VoiceChannel *discordgo.Channel
	TextChannel  *discordgo.Channel
	User         *discordgo.User
	Args         []string
	Sessions     *SessionManager
	Youtube      *Youtube
}

func NewContext(discord *discordgo.Session, guild *discordgo.Guild, textChannel *discordgo.Channel, user *discordgo.User, sessions *SessionManager, youtube *Youtube,
) *Context {
	ctx := new(Context)
	ctx.Discord = discord
	ctx.Guild = guild
	ctx.TextChannel = textChannel
	ctx.User = user
	ctx.Sessions = sessions
	ctx.Youtube = youtube
	return ctx
}

func (ctx *Context) Reply(content string) *discordgo.Message {
	msg, err := ctx.Discord.ChannelMessageSend(ctx.TextChannel.ID, content)
	if err != nil {
		fmt.Println("Error whilst sending message,", err)
		return nil
	}
	return msg
}

func (ctx *Context) GetVoiceChannel() *discordgo.Channel {
	if ctx.VoiceChannel != nil {
		return ctx.VoiceChannel
	}
	for _, state := range ctx.Guild.VoiceStates {
		if state.UserID == ctx.User.ID {
			channel, _ := ctx.Discord.State.Channel(state.ChannelID)
			ctx.VoiceChannel = channel
			return channel
		}
	}
	return nil
}

func (ctx *Context) CreateYoutubeEmbed() *[]*discordgo.MessageEmbed {
	thumbnail := &discordgo.MessageEmbedThumbnail{
		URL: ctx.Youtube.Search.Thumbnail.Url,
	}
	footer := &discordgo.MessageEmbedFooter{
		Text:    fmt.Sprintln("requested by " + ctx.User.String()),
		IconURL: ctx.User.AvatarURL(""),
	}
	embed := &discordgo.MessageEmbed{
		Title:       ctx.Youtube.Search.Title,
		Color:       0x142837,
		URL:         fmt.Sprintln("https://www.youtube.com/watch?v=" + ctx.Youtube.Search.Id),
		Description: ctx.Youtube.Search.Description,
		Thumbnail:   thumbnail,
		Footer:      footer,
	}
	embeds := []*discordgo.MessageEmbed{
		embed,
	}
	return &embeds
}

func (ctx *Context) CreateYoutubePlaylistEmbed(data string, url string) *[]*discordgo.MessageEmbed {
	// extract json data in embed struct
	var embed Embed
	err := json.Unmarshal([]byte(data), &embed)
	if err != nil {
		ctx.Reply(fmt.Sprint("Error whilst unmarshalling json," + err.Error()))
		return nil
	}

	thumbnail := &discordgo.MessageEmbedThumbnail{
		URL: embed.PlaylistThumbnail[0].Url,
	}
	footer := &discordgo.MessageEmbedFooter{
		Text:    fmt.Sprintln("requested by " + ctx.User.String()),
		IconURL: ctx.User.AvatarURL(""),
	}
	finalEmbed := &discordgo.MessageEmbed{
		Title:       embed.PlaylistTitle,
		Color:       0x142837,
		URL:         url,
		Description: "Playlist by " + embed.PlaylistUploader + "\n" + fmt.Sprintln("Total videos: "+fmt.Sprint(embed.PlayListCount)),
		Thumbnail:   thumbnail,
		Footer:      footer,
	}
	embeds := []*discordgo.MessageEmbed{
		finalEmbed,
	}
	return &embeds
}
