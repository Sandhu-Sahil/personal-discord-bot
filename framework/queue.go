package framework

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
