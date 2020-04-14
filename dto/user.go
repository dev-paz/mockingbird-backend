package dto

//User struct does a thing
type User struct {
	ID              string
	FirebaseUserID  string
	ProfilePhotoURL string
	Email           string
	Name            string
}

type CreateAccountRequest struct {
	FirebaseIDTokenID string `json:"firebase_id_token"`
	FirebaseUserID    string `json:"firebase_user_id"`
}
