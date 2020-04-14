package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	guuid "github.com/google/uuid"
	"github.com/mockingbird-backend/dto"
	"github.com/mockingbird-backend/models"
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

	// check whether user already exists
	existingUser, _ := models.ReadUser(createAccountRequest.FirebaseUserID)
	if existingUser != nil {
		fmt.Println("user exists, returning")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
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
	user := dto.User{
		ID:              "user_" + guuid.New().String(),
		FirebaseUserID:  firebaseIDToken.UID,
		Email:           fmt.Sprintf("%v", firebaseUser["email"]),
		Name:            fmt.Sprintf("%v", firebaseUser["name"]),
		ProfilePhotoURL: fmt.Sprintf("%v", firebaseUser["picture"]),
	}
	fmt.Println(user)
	err = models.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
