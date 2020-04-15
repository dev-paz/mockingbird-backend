package dto

//Song struct does a thing
type Song struct {
	Title      string `json:"title"`
	ID         string `json:"id"`
	Parts      string `json:"parts"`
	Difficulty int64  `json:"difficulty"`
}

type SongPart struct {
	ID       string `json:"id"`
	Part     int64  `json:"part"`
	MusicURL string `json:"music_url"`
	SongID   string `json:"song_id"`
}
