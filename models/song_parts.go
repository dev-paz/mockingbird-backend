package models

import (
	"fmt"

	"github.com/mockingbird-backend/dto"
)

func ReadSongParts(songID string) (*[]dto.SongPart, error) {
	var p dto.SongPart
	var parts []dto.SongPart
	query := `SELECT id, song_id, part, music_url FROM song_parts WHERE song_id=` + songID + `;`
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	for rows.Next() {
		err = rows.Scan(&p.ID, &p.SongID, &p.Part, &p.MusicURL)
		if err != nil {
			fmt.Println(err.Error())
			// handle this error
			panic(err)
		}
		parts = append(parts, p)
	}
	return &parts, nil
}

func CreateSongParts(sp []dto.SongPart) error {
	sqlStatement := `
	INSERT INTO song_parts (id, song_id, part, music_url)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	for _, p := range sp {
		id := ""
		err := db.QueryRow(sqlStatement, p.ID, p.SongID, p.Part, p.MusicURL).Scan(&id)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}
