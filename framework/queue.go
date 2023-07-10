package framework

import (
	"encoding/json"
)

type SongQueue struct {
	list    []Song
	current *Song
	Running bool
}

func newSongQueue() *SongQueue {
	queue := new(SongQueue)
	queue.list = make([]Song, 0)
	return queue
}

func (queue *SongQueue) Add(song Song) {
	queue.list = append(queue.list, song)
}

func (queue *SongQueue) HasNext() bool {
	return len(queue.list) > 0
}

func (queue *SongQueue) Next() Song {
	song := queue.list[0]
	queue.list = queue.list[1:]
	queue.current = &song
	return song
}

func (queue *SongQueue) Start(sess *Session, callback func(string)) {
	queue.Running = true
	for queue.HasNext() && queue.Running {
		song := queue.Next()
		callback("Now playing `" + song.Title + "`.")
		err := sess.Play(song)
		if err != nil {
			callback("Error playing song `" + song.Title + "`.")
			callback(err.Error())
			continue
		}
	}
	if !queue.Running {
		callback("Stopped playing.")
		queue.Running = false
	} else {
		callback("Finished queue.")
		queue.Running = false
	}
}

func (queue *SongQueue) StartPlaylist(data []string, ctx *Context, sess *Session, callback func(string)) {
	for index, line := range data {
		if len(line) == 0 {
			ctx.Reply("Panic, no songs found at a point in the playlist, recovered...")
			return
		}
		if line == "" || line == "\n" {
			continue
		}
		var video PlaylistVideo
		err := json.Unmarshal([]byte(line), &video)
		if err != nil {
			ctx.Reply("Panic, reading json: " + err.Error() + ", recovered...")
			continue
		}

		ctx.Youtube.Search.Id = video.Id
		types, outputTemp, err := ctx.Youtube.GetFromYT()
		if err != nil {
			ctx.Reply("Panic, song extraction error: " + err.Error() + ", recovered...")
		}
		if types != VIDEO_TYPE {
			ctx.Reply("A search in the playlist is not a video type or is not available at current time")
			continue
		}

		mainVideo, err := ctx.Youtube.Video(*outputTemp)
		if err != nil {
			ctx.Reply("Panic, reading json: " + err.Error() + ", recovered...")
			continue
		}
		song := NewSong(mainVideo.Media, mainVideo.Title, ctx.Youtube.Search.Id)
		sess.Queue.Add(*song)
		if index == 0 {
			if !sess.Queue.Running {
				go sess.Queue.Start(sess, func(msg string) {
					ctx.Reply(msg)
				})
			}
		}
	}
}

func (queue *SongQueue) Current() *Song {
	return queue.current
}

func (queue *SongQueue) Clear() {
	queue.list = make([]Song, 0)
}

func (queue *SongQueue) List() []Song {
	return queue.list
}

func (queue *SongQueue) Length() int {
	return len(queue.list)
}

func (queue *SongQueue) Remove(index int) string {
	song := queue.list[index]
	queue.list = append(queue.list[:index], queue.list[index+1:]...)
	if len(song.Title) < 50 {
		return song.Title
	}
	return song.Title[:50]
}
