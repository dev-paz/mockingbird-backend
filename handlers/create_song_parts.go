package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	guuid "github.com/google/uuid"
	"github.com/mockingbird-backend/dto"
	"github.com/mockingbird-backend/models"
)

func handleCreateSongParts(w http.ResponseWriter, req *http.Request) {
	songParts := []dto.SongPart{}

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&songParts)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	fmt.Println(songParts)

	for _, song := range songParts {
		song.ID = "sp_" + guuid.New().String()
	}

	err = models.CreateSongParts(songParts)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
