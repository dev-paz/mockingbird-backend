package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	guuid "github.com/google/uuid"

	"github.com/mockingbird-backend/dto"
	"github.com/mockingbird-backend/models"
)

func handleCreateProject(w http.ResponseWriter, req *http.Request) {
	createProjReq := dto.CreateProjectRequest{}

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&createProjReq)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	fmt.Println(createProjReq)

	openShotResp, err := createOpenShotProject()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	project := dto.ProjectDB{
		ID:          "proj_" + guuid.New().String(),
		Song:        createProjReq.SongID,
		Name:        createProjReq.Name,
		Created:     "time",
		Status:      "started",
		OpenshotURL: openShotResp.URL,
		OpenshotID:  fmt.Sprintf("%v", openShotResp.ID),
	}

	err = models.CreateProject(&project)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	clips := []dto.Clip{}
	for partID, userID := range createProjReq.ClipsToUsers {
		clip := dto.Clip{
			ID:                 "clip_" + guuid.New().String(),
			ProjectID:          project.ID,
			UserID:             userID,
			PartID:             partID,
			OpenshotProjectID:  project.OpenshotID,
			OpenshotProjectURL: project.OpenshotURL,
		}
		clips = append(clips, clip)
	}

	err = models.CreateClips(clips)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	resp, err := json.Marshal(&project)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

// creates a project on the openshot api
func createOpenShotProject() (dto.CreateProjectResponse, error) {
	data := dto.CreateProjectResponse{}
	username := "cloud-admin"
	passwd := "UZTWLVEr6n"

	requestData := []byte(`{
		"name": "Example lakweufiluqwh",
		"width": 1920,
		"height": 1080,
		"fps_num": 30,
		"fps_den": 1,
		"sample_rate": 44100,
		"channels": 2,
		"channel_layout": 3,
		"json": {}
}`)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://"+OpenShotIP+"/projects/", bytes.NewBuffer(requestData))
	if err != nil {
		return data, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, passwd)

	resp, err := client.Do(req)
	if err != nil {
		return data, err
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}
