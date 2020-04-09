package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mockingbird-backend/handlers"
	"github.com/mockingbird-backend/models"
)

// type data struct {
// 	URL    string `json:"url"`
// 	ID     int64  `json:"id"`
// 	Name   string `json:"name"`
// 	Height int64  `json:"height"`
// 	Width  int64  `json:"width"`
// }
//
// var username string = "cloud-admin"
// var passwd string = "UZTWLVEr6n"

const (
	host     = "localhost"
	port     = 5432
	user     = "paz"
	password = "password"
	dbname   = "mockingbird"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	models.InitDB(psqlInfo)

	handler.HandleRequests()
	log.Fatal(http.ListenAndServe(":8881", nil))
}

// func createProject() {
// 	data := data{}
//
// 	requestData := []byte(`{
// 		"name": "Example Project21345",
// 		"width": 1920,
// 		"height": 1080,
// 		"fps_num": 30,
// 		"fps_den": 1,
// 		"sample_rate": 44100,	createClip
//
// 		"channels": 2,
// 		"channel_layout": 3,
// 		"json": {}
// }`)
//
// 	client := &http.Client{}
// 	req, err := http.NewRequest("POST", "http://3.81.225.77/projects/", bytes.NewBuffer(requestData))
// 	req.Header.Set("X-Custom-Header", "myvalue")
// 	req.Header.Set("Content-Type", "application/json")
// 	req.SetBasicAuth(username, passwd)
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("error")
// 		log.Fatal(err)
// 	}
// 	decoder := json.NewDecoder(resp.Body)
// 	err = decoder.Decode(&data)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		panic(err)
// 	}
// }
//
// func deleteProject() {
// 	data := data{}
//
// 	client := &http.Client{}
// 	req, err := http.NewRequest("DELETE", "http://3.81.225.77/projects/7/", nil)
// 	req.SetBasicAuth(username, passwd)
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("error")
// 		log.Fatal(err)
// 	}
// 	decoder := json.NewDecoder(resp.Body)
// 	err = decoder.Decode(&data)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		panic(err)
// 	}
// }
//
// func createClip() {
// 	data := data{}
//
// 	requestData := []byte(`{
//     "file": "http://3.81.225.77/files/5/",
//     "position": 0,
//     "start": 0,
//     "end": 10,
//     "layer": 1,
// 		"project": "http://3.81.225.77/projects/8/",
//     "json": {
// 			"rotation": {
// 					"Points": [
// 							{
// 									"co": {
// 											"X": 1.0,
// 											"Y": 180.0
// 									},
// 									"interpolation": 2
// 							}
// 					]
// 			}
// 		}
// }`)
//
// 	client := &http.Client{}
// 	req, err := http.NewRequest("POST", "http://3.81.225.77/projects/8/clips/", bytes.NewBuffer(requestData))
// 	req.Header.Set("X-Custom-Header", "myvalue")
// 	req.Header.Set("Content-Type", "application/json")
// 	req.SetBasicAuth(username, passwd)
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("error")
// 		log.Fatal(err)
// 	}
// 	decoder := json.NewDecoder(resp.Body)
// 	err = decoder.Decode(&data)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		panic(err)
// 	}
// }
//
// func exportProject() {
// 	data := data{}
//
// 	requestData := []byte(`{
//     "export_type": "video",
//     "video_format": "mp4",
//     "video_codec": "libx264",
//     "video_bitrate": 8000000,
//     "audio_codec": "libfdk_aac",
//     "audio_bitrate": 1920000,
//     "start_frame": 1,
//     "end_frame": 100,
//     "project": "http://3.81.225.77/projects/8/",
//     "webhook": "",
//     "json": {},
//     "status": "pending"
// }`)
//
// 	client := &http.Client{}
// 	req, err := http.NewRequest("POST", "http://3.81.225.77/projects/8/exports/", bytes.NewBuffer(requestData))
// 	req.Header.Set("Content-Type", "application/json")
// 	req.SetBasicAuth(username, passwd)
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("error")
// 		log.Fatal(err)
// 	}
// 	decoder := json.NewDecoder(resp.Body)
// 	err = decoder.Decode(&data)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		panic(err)
// 	}
// }
