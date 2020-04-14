package handler

import (
	"net/http"
)

//HandleRequests routes incoming requests
func HandleRequests() {
	http.HandleFunc("/create_project", handleCreateProject)
	http.HandleFunc("/create_clip", handleUpdateClip)
	http.HandleFunc("/get_projects", handleGetProjects)
	http.HandleFunc("/get_songs", handleGetSongs)
	http.HandleFunc("/create_song", handleCreateSong)
	http.HandleFunc("/create_account", handleCreateAccount)
}

//// TODO: secure the routes by checking the firebase token

var OpenShotIP = "52.91.155.8"
