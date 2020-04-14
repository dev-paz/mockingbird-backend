package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/mockingbird-backend/dto"
	"google.golang.org/api/option"
)

func hangleSignIn(w http.ResponseWriter, req *http.Request) {

	signInRequest := dto.SignInRequest{}

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&signInRequest)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	_, err = client.VerifyIDToken(context.Background(), signInRequest.TokenID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
