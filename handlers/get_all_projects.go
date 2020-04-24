package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mockingbird-backend/models"
)

func handleGetAllProjects(w http.ResponseWriter, req *http.Request) {
	params, ok := req.URL.Query()["user_id"]
	if !ok {
		fmt.Println("Error parsing url query")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userID := params[0]
	if userID == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Missing param user id")
		return
	}

	projects, err := models.ReadAllProjects(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error reading from db")
		fmt.Println(err.Error())
		panic(err)
	}
	fmt.Println(projects)

	resp, err := json.Marshal(&projects)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		fmt.Println("Error marshalling")
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
