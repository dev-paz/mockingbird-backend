package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	handler "github.com/mockingbird-backend/handlers"
	"github.com/mockingbird-backend/models"
)

func main() {

	port := os.Getenv("PORT")

	models.InitDB()

	details, err := models.ReadOpenshotDetails()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	handler.OpenShotIP = details.IP

	handler.HandleRequests()
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "paz"
// 	password = "password"
// 	dbname   = "mockingbirdheroku"
// )
//
// func main() {
//
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)
// 	models.InitDB(psqlInfo)
//
// 	handler.HandleRequests()
// 	log.Fatal(http.ListenAndServe(":8881", nil))
// }
