package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/mockingbird-backend/dto"
	"github.com/mockingbird-backend/models"
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

	requestData := map[string]interface{}{
		"file":     "",
		"position": 0.0,
		"start":    0.0,
		"end":      project.Song.LengthSeconds,
		"layer":    1,
		"project":  "http://" + OpenShotIP + "/projects/" + project.OpenshotID + "/",
		"json": map[string]interface{}{
			"media_type": "audio",
		},
	}
	songParts, err := models.ReadSongParts(project.Song.ID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(project.Song.BackingTrack)
	if project.Song.BackingTrack != "" {
		addBackingTrack(project.Song.BackingTrack, OpenShotIP, requestData, 4, project.OpenshotID)
	}

	for layer, clip := range project.Clips {
		configData, err := fetchSongConfig((*songParts)[layer])
		if err != nil {
			panic(err)
		}

		configMap := map[string]interface{}{}
		json.Unmarshal(configData, &configMap)

		requestData["json"] = configMap["json"]
		requestData["file"] = clip.File
		requestData["layer"] = layer

		songConfig, err := json.Marshal(requestData)
		if err != nil {
			fmt.Println(err.Error())
		}

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

	exportResp, err := exportProject(project.ID, project.Song.ID, project.OpenshotID, project.Song.LengthSeconds)
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

func fetchSongConfig(songPart dto.SongPart) ([]byte, error) {
	fmt.Println(songPart.Config)
	resp, err := http.Get(songPart.Config)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err.Error())
		return nil, err
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
	}
	return content, nil
}

func exportProject(projectID string, songID string, openshotID string, length int64) (dto.ExportProjectResponse, error) {
	expResp := dto.ExportProjectResponse{}

	projectJSON := fmt.Sprintf(`{"ProjectID": "%s","SongID": "%s"}`, projectID, songID)
	data := map[string]interface{}{
		"export_type":   "video",
		"video_format":  "mp4",
		"video_codec":   "libx264",
		"video_bitrate": 8000000,
		"audio_codec":   "libfdk_aac",
		"audio_bitrate": 1920000,
		"start_frame":   1,
		"end_frame":     length * 30,
		"project":       "http://" + OpenShotIP + "/projects/" + openshotID + "/",
		"webhook":       "https://mockingbird-backend.herokuapp.com/create_music_video",
		"json": map[string]string{
			"hostname":  "ip-172-31-42-80",
			projectJSON: "success",
		},
	}

	requestData, err := json.Marshal(data)
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

func addBackingTrack(backingURL string, OpenShotIP string, requestData map[string]interface{}, layer int, OpenShotID string) error {
	client := http.Client{}

	out, err := os.Create("output.mp3")
	defer out.Close()
	resp, err := http.Get(backingURL)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("file downloaded")
	fmt.Println(resp.Body)

	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)

	b, w := createMultipartFormData("media", "output.mp3", OpenShotID)
	req, err := http.NewRequest("POST", "http://"+OpenShotIP+"/projects/"+OpenShotID+"/files/", &b)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	req.SetBasicAuth(username, passwd)
	fmt.Println("file uploaded")

	resp, err = client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	uploadResponse := struct {
		URL     string `json:"url"`
		ID      int64  `json:"id"`
		Project string `json:"project"`
	}{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&uploadResponse)
	if err != nil {
		fmt.Println(uploadResponse)
		fmt.Println(err.Error())
		panic(err)
	}

	jsonFile, err := os.Open("song_edit.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	configMap := map[string]interface{}{}
	json.Unmarshal(byteValue, &configMap)

	requestData["json"] = configMap["json"]
	requestData["file"] = uploadResponse.URL
	requestData["layer"] = layer

	songConfig, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println(err.Error())
	}

	req, err = http.NewRequest("POST", uploadResponse.Project+"clips/", bytes.NewBuffer(songConfig))
	if err != nil {
		fmt.Println(err.Error())
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, passwd)

	resp, err = client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
