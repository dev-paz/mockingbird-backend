package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handleGetIPAddress(w http.ResponseWriter, req *http.Request) {
	checkOpenshotIP()

	resp, err := json.Marshal(OpenShotIP)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Error marshalling")
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
