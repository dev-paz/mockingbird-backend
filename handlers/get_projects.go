package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mockingbird-backend/models"
)

func handleGetProjects(w http.ResponseWriter, req *http.Request) {

	projects, err := models.ReadProjectsforUser()
	if err != nil {
		fmt.Println("Error reading from db")
		fmt.Println(err.Error())
		panic(err)
	}

	resp, err := json.Marshal(&projects)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Error marshalling")
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
