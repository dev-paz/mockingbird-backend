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

func handleSignIn(w http.ResponseWriter, req *http.Request) {

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

	result, err := client.VerifyIDToken(context.Background(), signInRequest.TokenID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("checked token successfully")
	fmt.Println(result.UID)
	fmt.Println(result)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
