package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	guuid "github.com/google/uuid"
	"github.com/mockingbird-backend/dto"
	"github.com/mockingbird-backend/models"
)

func handleCreateSong(w http.ResponseWriter, req *http.Request) {
	song := dto.Song{}

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&song)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	song.ID = "song_" + guuid.New().String()

	err = models.CreateSong(&song)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
