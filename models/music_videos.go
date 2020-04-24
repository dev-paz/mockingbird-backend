package models

import (
	"fmt"

	"github.com/mockingbird-backend/dto"
)

func CreateMusicVideo(mv *dto.MusicVideo) error {
	sqlStatement := `
	INSERT INTO music_videos (id, url, created, song_id)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	id := ""
	err := db.QueryRow(sqlStatement, mv.ID, mv.URL, mv.Created, mv.SongID).Scan(&id)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func ReadMusicVideos() (*[]dto.MusicVideo, error) {
	var mv dto.MusicVideo
	var musicVideos []dto.MusicVideo
	rows, err := db.Query(`SELECT id, url, created, song_id FROM music_videos`)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	for rows.Next() {
		err = rows.Scan(&mv.ID, &mv.URL, &mv.Created, &mv.SongID)
		if err != nil {
			fmt.Println(err.Error())
			// handle this error
			panic(err)
		}
		musicVideos = append(musicVideos, mv)
	}
	return &musicVideos, nil
}
