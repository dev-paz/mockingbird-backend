package handler

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/mockingbird-backend/models"
)

func checkOpenshotIP() {
	details, err := models.ReadOpenshotDetails()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	OpenShotIP = details.IP
}

func createMultipartFormData(fieldName, fileName string, OpenShotID string) (bytes.Buffer, *multipart.Writer) {
	var b bytes.Buffer
	var err error
	w := multipart.NewWriter(&b)
	w.WriteField("project", "http://"+OpenShotIP+"/projects/"+OpenShotID+"/")
	w.WriteField("json", "{}")
	var fw io.Writer
	file := mustOpen(fileName)
	if fw, err = w.CreateFormFile(fieldName, file.Name()); err != nil {
		panic(err)
	}
	if _, err = io.Copy(fw, file); err != nil {
		panic(err)
	}
	w.Close()
	return b, w
}

func mustOpen(f string) *os.File {
	r, err := os.Open(f)
	if err != nil {
		pwd, _ := os.Getwd()
		fmt.Println("PWD: ", pwd)
		panic(err)
	}
	return r
}
