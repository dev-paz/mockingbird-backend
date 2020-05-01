package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mockingbird-backend/models"

	"github.com/mockingbird-backend/dto"
)

func handleDeleteProject(w http.ResponseWriter, req *http.Request) {
	checkOpenshotIP()

	delProjectReq := dto.DeleteProjectRequest{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&delProjectReq)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	project, err := models.ReadProject(delProjectReq.ProjectID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = deleteOpenShotProject(project.OpenshotID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = models.DeleteProject(project.ID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = models.DeleteClips(project.ID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// creates a project on the openshot api
func deleteOpenShotProject(openshotID string) error {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", "http://"+OpenShotIP+"/projects/"+openshotID+"/", nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, passwd)

	_, err = client.Do(req)
	if err != nil {
		return err
	}
	return nil
}
