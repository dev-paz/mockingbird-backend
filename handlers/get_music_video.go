package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mockingbird-backend/models"
)

func handleGetMusicVideo(w http.ResponseWriter, req *http.Request) {
	params, ok := req.URL.Query()["id"]
	if !ok {
		fmt.Println("Error parsing url query")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	musicVideoID := params[0]
	if musicVideoID == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Missing param project id")
		return
	}

	musicVideo, err := models.ReadMusicVideo(musicVideoID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error reading from db")
		fmt.Println(err.Error())
		panic(err)
	}

	resp, err := json.Marshal(&musicVideo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		fmt.Println("Error marshalling")
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
