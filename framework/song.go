package framework

type Song struct {
	Media    string
	Title    string
	Duration *string
	Id       string
}

func NewSong(media, title, id string) *Song {
	song := new(Song)
	song.Media = media
	song.Title = title
	song.Id = id
	return song
}
