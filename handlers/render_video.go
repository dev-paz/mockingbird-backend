package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/mockingbird-backend/dto"
	"github.com/mockingbird-backend/models"
	"google.golang.org/api/option"
)

func handleRenderVideo(w http.ResponseWriter, req *http.Request) {
	checkOpenshotIP()

	project := dto.Project{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&project)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	fmt.Println(project)

	requestData := map[string]interface{}{
		"file":     "",
		"position": 0.0,
		"start":    0.0,
		"end":      120.0,
		"layer":    1,
		"project":  "http://" + OpenShotIP + "/projects/" + project.OpenshotID + "/",
		"json":     "{}",
	}

	for layer, clip := range project.Clips {
		configData, err := fetchSongConfig(clip, project.Song.ID)
		if err != nil {
			panic(err)
		}

		fmt.Println(clip)

		configMap := map[string]interface{}{}
		json.Unmarshal(configData, &configMap)

		requestData["json"] = configMap["json"]
		requestData["file"] = clip.File
		requestData["layer"] = layer

		songConfig, err := json.Marshal(requestData)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println("POST", "http://"+OpenShotIP+"/projects/"+project.OpenshotID+"/clips/")
		fmt.Println(songConfig)

		client := &http.Client{}
		req, err := http.NewRequest("POST",
			"http://"+OpenShotIP+"/projects/"+project.OpenshotID+"/clips/",
			bytes.NewBuffer(songConfig))
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

		if resp.StatusCode != 201 {
			fmt.Println(resp.StatusCode)
			fmt.Println("Something went wrong createing a clip")
			return
		}
	}

	exportResp, err := exportProject(project.OpenshotID)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = models.UpdateProjectStatus(project.ID, "rendering", fmt.Sprintf("%v", exportResp.ID))
	if err != nil {
		fmt.Println(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func fetchSongConfig(clip dto.Clip, projectID string) ([]byte, error) {
	ctx := context.Background()
	sa := option.WithCredentialsFile("credentials.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	defer client.Close()

	ref, err := client.Collection("song_config").Doc(projectID).Get(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	songFile := ref.Data()[clip.PartID].(string)
	resp, err := http.Get(songFile)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return content, nil
}

func exportProject(projectID string) (dto.ExportProjectResponse, error) {
	expResp := dto.ExportProjectResponse{}
	data := map[string]interface{}{
		"export_type":   "video",
		"video_format":  "mp4",
		"video_codec":   "libx264",
		"video_bitrate": 8000000,
		"audio_codec":   "libfdk_aac",
		"audio_bitrate": 1920000,
		"start_frame":   1,
		"end_frame":     3600,
		"project":       "http://" + OpenShotIP + "/projects/" + projectID + "/",
		"webhook":       "https://mockingbird-backend.herokuapp.com/create_music_video",
		"json": map[string]string{
			"hostname": "ip-172-31-42-80",
			"webhook":  projectID,
		},
	}

	requestData, err := json.Marshal(data)
	fmt.Println(string(requestData))

	client := &http.Client{}
	req, err := http.NewRequest("POST",
		"http://"+OpenShotIP+"/projects/"+projectID+"/exports/",
		bytes.NewBuffer(requestData))
	if err != nil {
		fmt.Println(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, passwd)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return expResp, err
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&expResp)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	return expResp, nil
}
