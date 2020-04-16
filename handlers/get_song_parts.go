package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mockingbird-backend/models"
)

func handleGetSongParts(w http.ResponseWriter, req *http.Request) {

	songID := req.URL.Query().Get("song_id")
	if songID != "" {
		fmt.Println("Missing param song id")
		return
	}

	fmt.Println(songID)

	songParts, err := models.ReadSongParts(songID)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	resp, err := json.Marshal(&songParts)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
