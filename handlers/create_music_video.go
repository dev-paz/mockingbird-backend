package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	firebase "firebase.google.com/go"
	guuid "github.com/google/uuid"
	"google.golang.org/api/option"

	"github.com/mockingbird-backend/dto"
	"github.com/mockingbird-backend/models"
)

func handleCreateMusicVideo(w http.ResponseWriter, req *http.Request) {
	createVideoReq := dto.CreateMusicVideoRequest{}
	bucketLocation := "gs://mockingbird-287ec.appspot.com"

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&createVideoReq)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	videoID := "mv_" + guuid.New().String()
	fileName := videoID + ".mp4"
	musicVideo := dto.MusicVideo{
		ID:      videoID,
		URL:     bucketLocation + "/" + fileName,
		SongID:  createVideoReq.ProjectData.SongID,
		Created: createVideoReq.Created,
		Status:  "uploading",
	}

	go downloadFileToFirebase(createVideoReq.OutputURL, bucketLocation, fileName, videoID)

	err = models.CreateMusicVideo(&musicVideo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func downloadFileToFirebase(openShotURL string, bucketLocation string, fileName string, videoID string) error {
	fmt.Println("starting download...")
	config := &firebase.Config{
		StorageBucket: "gs://mockingbird-287ec.appspot.com",
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
		return err
	}
	if err := wc.Close(); err != nil {
		return err
	}

	err = models.UpdateVideoStatus(videoID, "completed")
	if err != nil {
		fmt.Println("error updating video status")
	}
	return nil
}
