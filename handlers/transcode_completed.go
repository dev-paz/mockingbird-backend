package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strings"

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
	fmt.Println(reqBody.ObjectKey)

	_, file := path.Split(reqBody.ObjectKey)
	videoID := strings.Split(file, ".")[0]
	err = models.UpdateProjectStatusByVideoID(videoID, "completed")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(file)
	fmt.Println(videoID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
