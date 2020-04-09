package handler

import (
	"encoding/json"
	"net/http"

	"github.com/mockingbird-backend/models"
)

func handleGetSongs(w http.ResponseWriter, req *http.Request) {

	songs, err := models.ReadSongs()
	if err != nil {
		panic(err)
	}

	resp, err := json.Marshal(&songs)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
