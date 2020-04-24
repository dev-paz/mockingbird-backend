package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mockingbird-backend/models"
)

func handleGetMusicVideos(w http.ResponseWriter, req *http.Request) {
	musicVideos, err := models.ReadMusicVideos()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error reading from db")
		fmt.Println(err.Error())
		panic(err)
	}

	resp, err := json.Marshal(&musicVideos)
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
