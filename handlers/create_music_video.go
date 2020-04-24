package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	guuid "github.com/google/uuid"

	"github.com/mockingbird-backend/dto"
	"github.com/mockingbird-backend/models"
)

func handleCreateMusicVideo(w http.ResponseWriter, req *http.Request) {
	createVideoReq := dto.CreateMusicVideoRequest{}

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&createVideoReq)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	musicVideo := dto.MusicVideo{
		ID:      "mv_" + guuid.New().String(),
		URL:     createVideoReq.OutputURL,
		SongID:  createVideoReq.ProjectData.SongID,
		Created: createVideoReq.Created,
	}

	err = models.CreateMusicVideo(&musicVideo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
