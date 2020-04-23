package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func handleCreateMusicVideo(w http.ResponseWriter, req *http.Request) {
	jsonResp, _ := ioutil.ReadAll(req.Body)
	fmt.Println(string(jsonResp))
}
