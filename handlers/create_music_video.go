package handler

import (
	"context"
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
	createVideoReq := new(dto.CreateMusicVideoRequest)
	bucketLocation := "mockingbird-287ec.appspot.com"

	err := req.ParseForm()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(req.Form)

	decoder := schema.NewDecoder()
	err = decoder.Decode(createVideoReq, req.Form)
	if err != nil {
		fmt.Println(err.Error())
	}

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
	}

	err = models.UpdateProjectStatus(createVideoReq.ProjectData.ProjectID, "uploading", "")
	if err != nil {
		fmt.Println(err.Error())
	}

	downloadFileToFirebase(downloadURL, bucketLocation, fileName, videoID, createVideoReq.ProjectData.ProjectID)

	err = models.CreateMusicVideo(&musicVideo)
	if err != nil {
		fmt.Println(err.Error())
		return
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
	fmt.Println("made it here")
	ctx := context.Background()
	client, err := app.Storage(ctx)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatalln(err)
	}
	fmt.Println("made it here")

	bucket, err := client.Bucket(bucketLocation)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatalln(err)
	}
	fmt.Println("made it here")
	fmt.Println(openShotURL)

	resp, err := http.Get(openShotURL)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	defer resp.Body.Close()
	fmt.Println("made it here")

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
	fmt.Println("made it here")

	err = models.UpdateProjectStatus(projectID, "completed", "")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("made it here")

	err = models.UpdateVideoStatus(videoID, "completed")
	if err != nil {
		fmt.Println("error updating video status")
	}
	fmt.Println("made it here")

	return nil
}
