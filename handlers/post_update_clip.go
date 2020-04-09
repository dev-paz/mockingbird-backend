package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mockingbird-backend/dto"
	"github.com/mockingbird-backend/models"
)

func handleUpdateClip(w http.ResponseWriter, req *http.Request) {
	clip := dto.Clip{}

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&clip)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	err = models.UpdateClip(clip.ID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// TODO: check whether thsi was the last file to be uploaded

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
