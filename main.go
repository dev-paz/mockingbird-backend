package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mockingbird-backend/handlers"
	"github.com/mockingbird-backend/models"
)

func main() {


	port := os.Getenv("PORT")

	models.InitDB()

	handler.HandleRequests()
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
