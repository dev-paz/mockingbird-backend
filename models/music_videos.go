package models

import (
	"fmt"

	"github.com/mockingbird-backend/dto"
)

func CreateMusicVideo(mv *dto.MusicVideo) error {
	sqlStatement := `
	INSERT INTO music_videos (id, url, created, song_id, status)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id`
	id := ""
	err := db.QueryRow(sqlStatement, mv.ID, mv.URL, mv.Created, mv.SongID, mv.Status).Scan(&id)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func ReadMusicVideos() (*[]dto.MusicVideo, error) {
	var mv dto.MusicVideo
	var musicVideos []dto.MusicVideo
	rows, err := db.Query(`SELECT id, url, created, song_id, status,
		project_name, owner, title, owner_photo, owner_name, album_art, clips
		 FROM music_video_view`)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	for rows.Next() {
		err = rows.Scan(&mv.ID, &mv.URL, &mv.Created, &mv.SongID, &mv.Status,
			&mv.Name, &mv.Owner, &mv.Title, &mv.OwnerPhoto, &mv.OwnerName, &mv.AlbumArt, &mv.Clips)
		if err != nil {
			fmt.Println(err.Error())
			// handle this error
			panic(err)
		}
		musicVideos = append(musicVideos, mv)
	}
	return &musicVideos, nil
}

func UpdateVideoStatus(id string, status string) error {
	sqlStatement := `
	UPDATE music_videos
	SET status = $2
	WHERE id=$1`
	_, err := db.Exec(sqlStatement, id, status)
	if err != nil {
		panic(err)
	}
	return nil
}
