package dto

//Project struct does a thing
type Project struct {
	Song        string `json:"song"`
	Name        string `json:"name"`
	ID          string `json:"id"`
	Created     string `json:"created"`
	Status      string `json:"status"`
	OpenshotID  int64  `json:"openshot_id"`
	OpenshotURL string `json:"openshot_url"`
}

//CreateProjectRequest struct for sending data to openshot api
type CreateProjectRequest struct {
	SongID       string           `json:"song_id"`
	UsersToClips map[string]int64 `json:"users_to_clips"`
}

//CreateProjectResponse struct for sending data to openshot api
type CreateProjectResponse struct {
	URL  string `json:"url"`
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
