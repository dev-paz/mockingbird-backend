package dto

type MusicVideo struct {
	ID         string `json:"id"`
	URL        string `json:"url"`
	Created    string `json:"created"`
	SongID     string `json:"song_id"`
	Status     string `json:"status"`
	Name       string `json:"name"`
	Owner      string `json:"owner"`
	Title      string `json:"title"`
	OwnerName  string `json:"owner_name"`
	OwnerPhoto string `json:"owner_photo"`
}

type CreateMusicVideoRequest struct {
	ExportURL   string      `json:"url" schema:"url"`
	ExportID    int64       `json:"id" schema:"id"`
	OutputURL   string      `json:"output" schema:"output"`
	ProjectURL  string      `json:"project" schema:"project"`
	Created     string      `json:"date_created" schema:"date_created"`
	ProjectData ProjectData `json:"json" schema:"json"`
}

type ProjectData struct {
	SongID    string
	ProjectID string
}
