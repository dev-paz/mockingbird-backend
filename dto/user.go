package dto

import jwt "github.com/dgrijalva/jwt-go"

//User struct does a thing
type User struct {
	ID              string
	GoogleUserID    string
	FirebaseUserID  string
	ProfilePhotoURL string
	GivenName       string
	FamilyName      string
	Email           string
	Name            string
}

type CreateAccountRequest struct {
	FirebaseIDTokenID string `json:"firebase_id_token"`
	GoogleIDTokenID   string `json:"google_id_token"`
}

// TokenInfo struct
type GoogleUser struct {
	Iss string `json:"iss"`
	// userId
	Sub string `json:"sub"`
	Azp string `json:"azp"`
	// clientId
	Aud string `json:"aud"`
	Iat int64  `json:"iat"`
	// expired time
	Exp int64 `json:"exp"`

	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	AtHash        string `json:"at_hash"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Local         string `json:"locale"`
	jwt.StandardClaims
}
