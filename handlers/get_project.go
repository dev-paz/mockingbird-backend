package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mockingbird-backend/models"
)

func handleGetProject(w http.ResponseWriter, req *http.Request) {

	params, ok := req.URL.Query()["id"]
	if !ok {
		fmt.Println("Error parsing url query")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(params)
	projectID := params[0]
	if projectID == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Missing param project id")
		return
	}

	projects, err := models.ReadProject(projectID)
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
