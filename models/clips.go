package models

import (
	"fmt"

	"github.com/mockingbird-backend/dto"

	_ "github.com/lib/pq"
)

func CreateClips(clips []dto.Clip) error {
	fmt.Println(clips)
	sqlStatement := `
	INSERT INTO clips (id, song_id, project_id, user_id, part, file, openshot_project_id, openshot_project_url)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING project_id`
	for _, c := range clips {
		id := ""
		err := db.QueryRow(sqlStatement, c.ID, c.SongID, c.ProjectID, c.UserID, c.Part, c.File, c.OpenshotProjectID, c.OpenshotProjectURL).Scan(&id)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}

func UpdateClip(id string) error {
	return nil
}
