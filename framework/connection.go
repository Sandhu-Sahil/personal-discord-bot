package framework

import (
	"sync"

	"github.com/bwmarrin/discordgo"
)

type Connection struct {
	voiceConnection *discordgo.VoiceConnection
	send            chan []int16
	lock            sync.Mutex
	sendpcm         bool
	stopRunning     bool
	playing         bool
	paused          bool
	loop            bool
	replay          bool
}

func NewConnection(voiceConnection *discordgo.VoiceConnection) *Connection {
	connection := new(Connection)
	connection.voiceConnection = voiceConnection
	return connection
}

func (connection *Connection) Disconnect() {
	connection.voiceConnection.Disconnect()
}

func (connection *Connection) Pause() {
	connection.paused = true
}

func (connection *Connection) Resume() {
	connection.paused = false
}

func (connection *Connection) ToogleLoop() {
	connection.loop = !connection.loop
}

func (connection *Connection) ToogleReplay() {
	connection.replay = !connection.replay
}
