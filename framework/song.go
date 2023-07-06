package framework

import (
	"os"
	"os/exec"
	"strconv"
)

type Song struct {
	Media    string
	Title    string
	Duration *string
	Id       string
}

func (song *Song) Ffmpeg() *exec.Cmd {
	// return exec.Command("ffmpeg", "-i", song.Media, "-f", "s16le", "-ar", strconv.Itoa(FRAME_RATE), "-ac",
	// 	strconv.Itoa(CHANNELS), "-probesize", "32M", "pipe:1")

	outputFileName := "imports/" + song.Id + ".mp4"

	// check if the file already exists
	if _, err := os.Stat(outputFileName); err == nil {
		// file exists, return the command
		return exec.Command("ffmpeg", "-i", outputFileName, "-f", "s16le", "-ar", strconv.Itoa(FRAME_RATE), "-ac", strconv.Itoa(CHANNELS), "pipe:1")
	}

	cmdimport := exec.Command("ffmpeg", "-y", "-i", song.Media, "-c", "copy", outputFileName)
	err := cmdimport.Run()
	if err != nil {
		return nil
	}

	cmd := exec.Command("ffmpeg", "-i", outputFileName, "-f", "s16le", "-ar", strconv.Itoa(FRAME_RATE), "-ac", strconv.Itoa(CHANNELS), "pipe:1")

	// // if already exists, delete the log file
	// if _, err := os.Stat("logs/ffmpeg.log"); err == nil {
	// 	os.Remove("logs/ffmpeg.log")
	// }
	// // Create a file to store the FFmpeg logs
	// logFile, _ := os.Create("logs/ffmpeg.log")

	// // Set the command's stderr to the log file
	// cmd.Stderr = logFile

	return cmd
}

func NewSong(media, title, id string) *Song {
	song := new(Song)
	song.Media = media
	song.Title = title
	song.Id = id
	return song
}
