package handler

import (
	"encoding/json"
	"net/http"

	"github.com/mockingbird-backend/models"
)

func handlegetUsers(w http.ResponseWriter, req *http.Request) {

	users, err := models.ReadAllUsers()
	if err != nil {
		panic(err)
	}

	resp, err := json.Marshal(&users)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
