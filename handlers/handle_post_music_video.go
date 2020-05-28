package handler

import (
	"fmt"
	"net/http"

	"github.com/mockingbird-backend/models"
)

func handlePostMusicVideo(w http.ResponseWriter, req *http.Request) {
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

	err := models.UpdateVideoPublic(musicVideoID, true)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
