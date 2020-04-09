package handler

import (
	"encoding/json"
	"net/http"

	"github.com/mockingbird-backend/models"
)

func handleGetProjects(w http.ResponseWriter, req *http.Request) {

	projects, err := models.ReadProjectsforUser()
	if err != nil {
		panic(err)
	}

	resp, err := json.Marshal(&projects)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
