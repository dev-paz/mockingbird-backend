package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	jwt "github.com/dgrijalva/jwt-go"
	guuid "github.com/google/uuid"
	"github.com/mockingbird-backend/dto"
	"google.golang.org/api/option"
)

func handleCreateAccount(w http.ResponseWriter, req *http.Request) {
	createAccountRequest := dto.CreateAccountRequest{}

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&createAccountRequest)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	opt := option.WithCredentialsFile("credentials.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	client, err := app.Auth(context.Background())
	if err != nil {
		fmt.Println(err.Error())
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	firebaseIDToken, err := client.VerifyIDToken(context.Background(), createAccountRequest.FirebaseIDTokenID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	firebaseUser := firebaseIDToken.Claims

	if createAccountRequest.GoogleIDTokenID != "" {
		googleUser, err := getGoogleUser(createAccountRequest.GoogleIDTokenID)
		fmt.Println(googleUser)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		_ = dto.User{
			ID:              "user_" + guuid.New().String(),
			Name:            googleUser.Name,
			FamilyName:      googleUser.FamilyName,
			GivenName:       googleUser.GivenName,
			ProfilePhotoURL: googleUser.Picture,
			Email:           fmt.Sprintf("%v", firebaseUser["poo"]),
			GoogleUserID:    googleUser.Sub,
			FirebaseUserID:  firebaseIDToken.UID,
		}
	}

	_ = dto.User{
		ID:              "user_" + guuid.New().String(),
		FirebaseUserID:  firebaseIDToken.UID,
		Email:           fmt.Sprintf("%v", firebaseUser["email"]),
		Name:            fmt.Sprintf("%v", firebaseUser["name"]),
		ProfilePhotoURL: fmt.Sprintf("%v", firebaseUser["picture"]),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func getGoogleUser(googleToken string) (*dto.GoogleUser, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(googleToken, &dto.GoogleUser{})
	if googleUser, ok := token.Claims.(*dto.GoogleUser); ok {
		return googleUser, nil
	}
	return nil, fmt.Errorf("error getting user %s", err)
}
