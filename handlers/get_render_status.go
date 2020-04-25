package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mockingbird-backend/dto"
)

func handleGetRenderStatus(w http.ResponseWriter, req *http.Request) {
	expResp := dto.ExportProjectResponse{}

	params, ok := req.URL.Query()["id"]
	if !ok {
		fmt.Println("Error parsing url query")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	openshotID := params[0]
	if openshotID == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Missing param openshotID")
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://"+OpenShotIP+"/exports/"+openshotID+"/", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, passwd)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&expResp)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	expRespJSON, err := json.Marshal(&expResp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		fmt.Println("Error marshalling")
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(expRespJSON)
}
