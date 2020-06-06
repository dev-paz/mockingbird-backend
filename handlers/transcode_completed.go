package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/mockingbird-backend/dto"
	"github.com/mockingbird-backend/models"
)

// update music video to completed

func handleTranscodeCompleted(w http.ResponseWriter, req *http.Request) {
	reqBody := struct {
		ObjectKey string `json:"objectKey"`
	}{}

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&reqBody)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	_, file := path.Split(reqBody.ObjectKey)
	videoID := strings.Split(file, ".")[0]

	project, err := models.ReadProjectByVideoID(videoID)
	if err != nil {
		fmt.Println(err.Error())
	}

	musicVideo := dto.MusicVideo{
		ID:      videoID,
		URL:     "https://" + reqBody.ObjectKey,
		SongID:  project.Song.ID,
		Created: project.Created,
		Status:  "uploading",
		Public:  false,
		Project: project.ID,
	}

	err = models.CreateMusicVideo(&musicVideo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = models.UpdateProjectStatus(project.ID, "completetd", project.ExportID)
	if err != nil {
		fmt.Println(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
