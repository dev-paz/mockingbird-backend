package models

import (
	"fmt"

	"github.com/mockingbird-backend/dto"
)

func CreateMusicVideo(mv *dto.MusicVideo) error {
	sqlStatement := `
	INSERT INTO projects (id, url, created, song_id)
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
