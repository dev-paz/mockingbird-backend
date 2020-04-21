package dto

//Project struct does a thing
type Project struct {
	Song        Song      `json:"song" db:"song"`
	Clips       ClipSlice `json:"clips" db:"clips"`
	Name        string    `json:"name" db:"name"`
	Users       UserSlice `json:"users" db:"users"`
	ID          string    `json:"id" db:"id"`
	Created     string    `json:"created" db:"created"`
	Status      string    `json:"status" db:"status"`
	OpenshotID  string    `json:"openshot_id" db:"openshot_id"`
	OpenshotURL string    `json:"openshot_url" db:"openshot_url"`
}

type ProjectDB struct {
	Song        string `json:"song" db:"song"`
	Name        string `json:"name" db:"name"`
	ID          string `json:"id" db:"id"`
	Created     string `json:"created" db:"created"`
	Status      string `json:"status" db:"status"`
	OpenshotID  string `json:"openshot_id" db:"openshot_id"`
	OpenshotURL string `json:"openshot_url" db:"openshot_url"`
}

//CreateProjectRequest struct for sending data to openshot api
type CreateProjectRequest struct {
	SongID       string            `json:"song_id"`
	ClipsToUsers map[string]string `json:"clips_to_users"`
	Name         string            `json:"name"`
}

//CreateProjectResponse struct for sending data to openshot api
type CreateProjectResponse struct {
	URL  string `json:"url"`
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type GetAllProjectsResponse struct {
	Projects []Project `json:"projects"`
}

type GetAllProjectsRequest struct {
	UserID string `json:"user_id"`
}
