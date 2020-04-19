package dto

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	_ "github.com/lib/pq"
)

//User struct does a thing
type User struct {
	ID              string `json:"id" db:"id"`
	FirebaseUserID  string `json:"firebase_user_id" db:"firebase_user_id"`
	ProfilePhotoURL string `json:"profile_photo_url" db:"profile_photo_url"`
	Email           string `json:"email" db:"email"`
	Name            string `json:"name" db:"name"`
}

type CreateAccountRequest struct {
	FirebaseIDTokenID string `json:"firebase_id_token"`
	FirebaseUserID    string `json:"firebase_user_id"`
}

type UserSlice []User

func (u User) Value() (driver.Value, error) {
	return json.Marshal(u)
}

func (u *UserSlice) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, u)
	case string:
		return json.Unmarshal([]byte(v), u)
	}
	return errors.New("type assertion failed")
}
