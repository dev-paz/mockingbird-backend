package dto

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

//Clip struct does a thing
type Clip struct {
	ID                 string `json:"id"`
	SongID             string `json:"song_id"`
	ProjectID          string `json:"project_id"`
	UserID             string `json:"user_id"`
	PartID             string `json:"part_id"`
	File               string `json:"file"`
	OpenshotProjectID  string `json:"openshot_project_id"`
	OpenshotProjectURL string `json:"openshot_project_url"`
	Status             string `json:"status"`
}

type UpdateClipRequest struct {
	ID   string `json:"id"`
	File string `json:"file"`
}

type ClipSlice []Clip

func (u Clip) Value() (driver.Value, error) {
	return json.Marshal(u)
}

func (u *ClipSlice) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, u)
	case string:
		return json.Unmarshal([]byte(v), u)
	}
	return errors.New("type assertion failed")
}
