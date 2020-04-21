package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mockingbird-backend/dto"
	"github.com/mockingbird-backend/models"
)

func handleUpdateClip(w http.ResponseWriter, req *http.Request) {
	clipsReq := dto.UpdateClipRequest{}

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&clipsReq)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	err = models.UpdateClip(clipsReq.ID, clipsReq.File)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
