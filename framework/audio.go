package framework

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"layeh.com/gopus"
)

const (
	CHANNELS   int = 2
	FRAME_RATE int = 48000
	FRAME_SIZE int = 960
	MAX_BYTES  int = (FRAME_SIZE * 2) * 2
)

/*
this shit is messy and i don't fully understand it yet credit to github.com/bwmarrin's voice example for the base code
*/

func (connection *Connection) sendPCM(voice *discordgo.VoiceConnection, pcm <-chan []int16) {
	connection.lock.Lock()
	if connection.sendpcm || pcm == nil {
		connection.lock.Unlock()
		return
	}
	connection.sendpcm = true
	connection.lock.Unlock()
	defer func() {
		connection.sendpcm = false
	}()
	encoder, err := gopus.NewEncoder(FRAME_RATE, CHANNELS, gopus.Audio)
	if err != nil {
		fmt.Println("NewEncoder error,", err)
		return
	}
	for {
		if connection.paused {
			continue
		}

		receive, ok := <-pcm
		if !ok {
			fmt.Println("PCM channel closed")
			return
		}
		opus, err := encoder.Encode(receive, FRAME_SIZE, MAX_BYTES)
		if err != nil {
			fmt.Println("Encoding error,", err)
			return
		}
		if !voice.Ready || voice.OpusSend == nil {
			fmt.Printf("Discordgo not ready for opus packets. %+v : %+v", voice.Ready, voice.OpusSend)
			return
		}
		voice.OpusSend <- opus
	}
}

func (connection *Connection) Play(ffmpeg *exec.Cmd) error {
	if connection.playing {
		return errors.New("song already playing")
	}
	connection.stopRunning = false
	out, err := ffmpeg.StdoutPipe()
	if err != nil {
		return err
	}
	buffer := bufio.NewReaderSize(out, 16384)
	err = ffmpeg.Start()
	if err != nil {
		return err
	}
	connection.playing = true
	defer func() {
		connection.playing = false
	}()
	connection.paused = false
	connection.voiceConnection.Speaking(true)
	defer connection.voiceConnection.Speaking(false)
	if connection.send == nil {
		connection.send = make(chan []int16, 2)
	}
	go connection.sendPCM(connection.voiceConnection, connection.send)
	for {
		// if disconnected
		if !connection.voiceConnection.Ready {
			connection.DeleteImportFile(ffmpeg.Args[2])
			ffmpeg.Process.Kill()
			break
		}
		if connection.stopRunning {
			connection.DeleteImportFile(ffmpeg.Args[2])
			ffmpeg.Process.Kill()
			break
		}

		// if paused
		if connection.paused {
			continue
		}

		// if replaying
		if connection.replay {
			connection.replay = false
			ffmpeg = exec.Command("ffmpeg", "-i", ffmpeg.Args[2], "-f", "s16le", "-ar", strconv.Itoa(FRAME_RATE), "-ac", strconv.Itoa(CHANNELS), "pipe:1")
			out, err = ffmpeg.StdoutPipe()
			if err != nil {
				return err
			}
			buffer = bufio.NewReaderSize(out, 16384)
			err = ffmpeg.Start()
			if err != nil {
				return err
			}
			continue
		}

		audioBuffer := make([]int16, FRAME_SIZE*CHANNELS)
		err = binary.Read(buffer, binary.LittleEndian, &audioBuffer)

		if err == io.EOF || err == io.ErrUnexpectedEOF {
			if connection.loop {
				// restart song
				ffmpeg = exec.Command("ffmpeg", "-i", ffmpeg.Args[2], "-f", "s16le", "-ar", strconv.Itoa(FRAME_RATE), "-ac", strconv.Itoa(CHANNELS), "pipe:1")
				out, err = ffmpeg.StdoutPipe()
				if err != nil {
					return err
				}
				buffer = bufio.NewReaderSize(out, 16384)
				err = ffmpeg.Start()
				if err != nil {
					return err
				}
				continue
			} else {
				connection.DeleteImportFile(ffmpeg.Args[2])
				return nil
			}
		}
		if err != nil {
			connection.DeleteImportFile(ffmpeg.Args[2])
			return err
		}
		connection.send <- audioBuffer
	}
	return nil
}

func (connection *Connection) DeleteImportFile(file string) {
	err := os.Remove(file)
	if err != nil {
		fmt.Println("Error deleting import file,", err)
	}
}

func (connection *Connection) Stop() {
	connection.stopRunning = true
	connection.playing = false
}
