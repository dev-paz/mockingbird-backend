package handler

import (
	"net/http"
)

//HandleRequests routes incoming requests
func HandleRequests() {
	http.HandleFunc("/get_ip", handleGetIPAddress)

	http.HandleFunc("/create_project", handleCreateProject)
	http.HandleFunc("/delete_project", handleDeleteProject)
	http.HandleFunc("/get_project", handleGetProject)
	http.HandleFunc("/get_all_projects", handleGetAllProjects)
	http.HandleFunc("/render_video", handleRenderVideo)
	http.HandleFunc("/get_render_status", handleGetRenderStatus)

	http.HandleFunc("/update_clip", handleUpdateClip)

	http.HandleFunc("/get_songs", handleGetSongs)
	http.HandleFunc("/create_song", handleCreateSong)

	http.HandleFunc("/create_song_parts", handleCreateSongParts)
	http.HandleFunc("/get_song_parts", handleGetSongParts)

	http.HandleFunc("/create_account", handleCreateAccount)
	http.HandleFunc("/get_all_users", handleGetAllUsers)

	http.HandleFunc("/create_music_video", handleCreateMusicVideo)
	http.HandleFunc("/get_music_videos", handleGetMusicVideos)
	http.HandleFunc("/post_music_video", handlePostMusicVideo)
	http.HandleFunc("/get_music_video", handleGetMusicVideo)

	http.HandleFunc("/upload_video", handleUploadVideoRequest)
}

// TODO: secure the routes by checking the firebase token

var OpenShotIP = ""
var username = "cloud-admin"
var passwd = "3lUkJFLq5g"
