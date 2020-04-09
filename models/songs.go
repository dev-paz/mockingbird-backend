package models

import (
	"fmt"

	"github.com/mockingbird-backend/dto"
)

func ReadSongs() (*[]dto.Song, error) {
	var s dto.Song
	var songs []dto.Song
	rows, err := db.Query(`SELECT id, title, parts FROM songs`)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	for rows.Next() {
		err = rows.Scan(&s.ID, &s.Name, &s.Parts)
		if err != nil {
			// handle this error
			panic(err)
		}
		songs = append(songs, s)
	}
	return &songs, nil
}
