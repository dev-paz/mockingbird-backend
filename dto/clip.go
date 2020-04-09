package dto

//Clip struct does a thing
type Clip struct {
	ID                 string `json:"id"`
	SongID             string `json:"song_id"`
	ProjectID          string `json:"project_id"`
	UserID             string `json:"user_id"`
	Part               int64  `json:"part"`
	File               string `json:"file"`
	OpenshotProjectID  int64  `json:"openshot_project_id"`
	OpenshotProjectURL string `json:"openshot_project_url"`
	Status             string `json:"status"`
}
