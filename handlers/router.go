package handler

import (
	"net/http"
)

//HandleRequests routes incoming requests
func HandleRequests() {
	http.HandleFunc("/create_project", handleCreateProject)
	http.HandleFunc("/get_project", handleGetProject)
	http.HandleFunc("/get_all_projects", handleGetAllProjects)
	http.HandleFunc("/render_video", handleRenderVideo)

	http.HandleFunc("/update_clip", handleUpdateClip)

	http.HandleFunc("/get_songs", handleGetSongs)
	http.HandleFunc("/create_song", handleCreateSong)

	http.HandleFunc("/create_song_parts", handleCreateSongParts)
	http.HandleFunc("/get_song_parts", handleGetSongParts)

	http.HandleFunc("/create_account", handleCreateAccount)
	http.HandleFunc("/get_all_users", handleGetAllUsers)

	http.HandleFunc("/create_music_video", handleCreateMusicVideo)
	http.HandleFunc("/get_music_videos", handleGetMusicVideos)

}

// TODO: secure the routes by checking the firebase token

var OpenShotIP = "3.92.160.40"
var username = "cloud-admin"
var passwd = "UZTWLVEr6n"
