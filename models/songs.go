package models

import (
	"fmt"

	"github.com/mockingbird-backend/dto"
)

func ReadSongs() (*[]dto.Song, error) {
	var s dto.Song
	var songs []dto.Song
	rows, err := db.Query(`SELECT id, title, parts, difficulty FROM songs`)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	for rows.Next() {
		err = rows.Scan(&s.ID, &s.Title, &s.Parts, &s.Difficulty)
		if err != nil {
			// handle this error
			panic(err)
		}
		songs = append(songs, s)
	}
	return &songs, nil
}

func CreateSong(s *dto.Song) error {
	sqlStatement := `
	INSERT INTO songs (id, title, parts, difficulty)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	id := ""
	err := db.QueryRow(sqlStatement, s.ID, s.Title, s.Parts, s.Difficulty).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}
