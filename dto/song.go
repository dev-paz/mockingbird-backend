package dto

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

//Song struct does a thing
type Song struct {
	Title         string `json:"title" db:"title"`
	ID            string `json:"id" db:"id"`
	Parts         int64  `json:"parts" db:"parts"`
	Difficulty    int64  `json:"difficulty" db:"difficulty"`
	LengthSeconds int64  `json:"length_seconds" db:"length_seconds"`
	BackingTrack  string `json:"backing_track" db:"backing_track"`
}

type SongPart struct {
	ID          string  `json:"id" db:"difficulty"`
	Part        int64   `json:"part" db:"part"`
	MusicURL    string  `json:"music_url" db:"music_url"`
	SongID      string  `json:"song_id" db:"song_id"`
	Type        string  `json:"type" db:"type"`
	AspectRatio float64 `json:"aspect_ratio" db:"aspect_ratio"`
	Config      string  `json:"config" db:"config"`
}

func (u Song) Value() (driver.Value, error) {
	return json.Marshal(u)
}

func (u *Song) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, u)
	case string:
		return json.Unmarshal([]byte(v), u)
	}
	return errors.New("type assertion failed")
}
