package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gorilla/schema"

	guuid "github.com/google/uuid"

	"github.com/mockingbird-backend/dto"
	"github.com/mockingbird-backend/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func handleCreateMusicVideo(w http.ResponseWriter, req *http.Request) {
	checkOpenshotIP()
	projectData := dto.ProjectData{}
	createVideoReq := new(dto.CreateMusicVideoRequest)

	err := req.ParseForm()
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, str := range req.Form["json"] {
		fmt.Println(str)
		if str == "hostname" || str == "status" {
			continue
		}
		json.Unmarshal([]byte(str), &projectData)
	}

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
		Public:  false,
		Project: projectData.ProjectID,
	}

	err = models.UpdateProjectMusicVideo(createVideoReq.ProjectData.ProjectID, videoID)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = downloadFileToS3(downloadURL, fileName)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = models.CreateMusicVideo(&musicVideo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func downloadFileToS3(openShotURL string, fileName string) error {
	resp, err := http.Get(openShotURL)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	const (
		AWS_S3_REGION = "us-east-1"
		AWS_S3_BUCKET = "mockingbird-source-16iqlqkop2oyn"
	)

	s, err := session.NewSession(&aws.Config{Region: aws.String(AWS_S3_REGION)})
	if err != nil {
		fmt.Println(err.Error())
	}

	r, err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket: aws.String(AWS_S3_BUCKET),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(bodyBytes),
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(r)
	return nil
}
