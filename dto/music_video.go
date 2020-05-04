package dto

type MusicVideo struct {
	ID      string `json:"id"`
	URL     string `json:"url"`
	Created string `json:"created"`
	SongID  string `json:"song_id"`
	Status  string `json:"status"`
}

type CreateMusicVideoRequest struct {
	ExportURL   string            `json:"url" schema:"url"`
	ExportID    int64             `json:"id" schema:"id"`
	OutputURL   string            `json:"output" schema:"output"`
	ProjectURL  string            `json:"project" schema:"project"`
	Created     string            `json:"date_created" schema:"date_created"`
	ProjectData map[string]string `json:"json" schema:"json"`
}

type ProjectData struct {
	SongID    string `json:"song_id"`
	ProjectID string `json:"project_id"`
}
