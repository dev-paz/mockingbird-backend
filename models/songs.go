package models

import (
	"fmt"

	"github.com/mockingbird-backend/dto"
)

func ReadSongs() (*[]dto.Song, error) {
	var s dto.Song
	var songs []dto.Song
	rows, err := db.Query(`SELECT id, title, parts, difficulty, length_seconds, backing_track FROM songs`)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	for rows.Next() {
		err = rows.Scan(&s.ID, &s.Title, &s.Parts, &s.Difficulty, &s.LengthSeconds, &s.BackingTrack)
		if err != nil {
			fmt.Println(err.Error())
			// handle this error
			panic(err)
		}
		songs = append(songs, s)
	}
	return &songs, nil
}

func CreateSong(s *dto.Song) error {
	sqlStatement := `
	INSERT INTO songs (id, title, parts, difficulty, length_seconds, backing_track)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id`
	id := ""
	err := db.QueryRow(sqlStatement, s.ID, s.Title, s.Parts, s.Difficulty, s.LengthSeconds, s.BackingTrack).Scan(&id)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
