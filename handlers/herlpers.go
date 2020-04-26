package handler

import (
	"fmt"

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
