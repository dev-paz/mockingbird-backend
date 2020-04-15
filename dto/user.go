package dto

//User struct does a thing
type User struct {
	ID              string `json:"id"`
	FirebaseUserID  string `json:"firebase_user_id"`
	ProfilePhotoURL string `json:"profile_photo_url"`
	Email           string `json:"email"`
	Name            string `json:"name"`
}

type CreateAccountRequest struct {
	FirebaseIDTokenID string `json:"firebase_id_token"`
	FirebaseUserID    string `json:"firebase_user_id"`
}
