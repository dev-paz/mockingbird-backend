package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/schema"

	firebase "firebase.google.com/go"
	guuid "github.com/google/uuid"
	"google.golang.org/api/option"

	"github.com/mockingbird-backend/dto"
	"github.com/mockingbird-backend/models"
)

func handleCreateMusicVideo(w http.ResponseWriter, req *http.Request) {
	checkOpenshotIP()
	projectData := dto.ProjectData{}
	createVideoReq := new(dto.CreateMusicVideoRequest)
	bucketLocation := "mockingbird-287ec.appspot.com"

	err := req.ParseForm()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(req.Form)

	for _, str := range req.Form["json"] {
		fmt.Println(str)
		if str == "hostname" || str == "status" {
			continue
		}
		json.Unmarshal([]byte(str), &projectData)
	}

	fmt.Println(projectData)
	createVideoReq.ProjectData = projectData

	decoder := schema.NewDecoder()
	err = decoder.Decode(createVideoReq, req.Form)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(createVideoReq)
	fmt.Println(createVideoReq.ProjectData)

	u, err := url.Parse(createVideoReq.OutputURL)
	if err != nil {
		panic(err)
	}
	downloadURL := "http://" + OpenShotIP + u.Path

	videoID := "mv_" + guuid.New().String()
	fileName := videoID + ".mp4"
	musicVideo := dto.MusicVideo{
		ID:      videoID,
		URL:     fileName,
		SongID:  createVideoReq.ProjectData.SongID,
		Created: createVideoReq.Created,
		Status:  "uploading",
		Public:  false,
		Project: projectData.ProjectID,
	}

	downloadFileToFirebase(downloadURL, bucketLocation, fileName, videoID, createVideoReq.ProjectData.ProjectID)

	err = models.CreateMusicVideo(&musicVideo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = models.UpdateProjectStatus(createVideoReq.ProjectData.ProjectID, "completed", string(createVideoReq.ExportID))
	if err != nil {
		fmt.Println(err.Error())
	}

	err = models.UpdateProjectMusicVideo(createVideoReq.ProjectData.ProjectID, videoID)
	if err != nil {
		fmt.Println(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func downloadFileToFirebase(openShotURL string, bucketLocation string, fileName string, videoID string, projectID string) error {
	fmt.Println("starting download...")
	config := &firebase.Config{
		StorageBucket: "mockingbird-287ec.appspot.com",
	}
	opt := option.WithCredentialsFile("credentials.json")
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatalln(err)
	}
	ctx := context.Background()
	client, err := app.Storage(ctx)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatalln(err)
	}

	bucket, err := client.Bucket(bucketLocation)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatalln(err)
	}

	resp, err := http.Get(openShotURL)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	defer resp.Body.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()
	wc := bucket.Object(fileName).NewWriter(ctx)
	if _, err = io.Copy(wc, resp.Body); err != nil {
		fmt.Println(err.Error())
		return err
	}
	if err := wc.Close(); err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = models.UpdateVideoStatus(videoID, "completed")
	if err != nil {
		fmt.Println("error updating video status")
	}

	return nil
}
