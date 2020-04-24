package dto

type MusicVideo struct {
	ID      string `json:"id"`
	URL     string `json:"url"`
	Created string `json:"created"`
	SongID  string `json:"song_id"`
	Status  string `json:"status"`
}

type CreateMusicVideoRequest struct {
	ExportURL   string      `json:"url"`
	ExportID    int64       `json:"id"`
	OutputURL   string      `json:"output"`
	ProjectURL  string      `json:"project"`
	Created     string      `json:"date_created"`
	ProjectData ProjectData `json:"json"`
}

type ProjectData struct {
	SongID    string `json:"song_id"`
	ProjectID string `json:"project_id"`
}
