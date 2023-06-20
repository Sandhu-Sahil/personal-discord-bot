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
