package models

import (
	"fmt"

	"github.com/mockingbird-backend/dto"

	_ "github.com/lib/pq"
)

func CreateClips(clips []dto.Clip) error {
	fmt.Println("printing clips")
	fmt.Println(clips)
	sqlStatement := `
	INSERT INTO clips (id, song_id, project_id, user_id, part_id, file, openshot_project_id, openshot_project_url)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING project_id`
	for _, c := range clips {
		id := ""
		err := db.QueryRow(sqlStatement, c.ID, c.SongID, c.ProjectID, c.UserID, c.PartID, c.File, c.OpenshotProjectID, c.OpenshotProjectURL).Scan(&id)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}

func UpdateClip(id string, file string) error {
	sqlStatement := `
	UPDATE clips
	SET file = $2
	WHERE id=$1`

	_, err := db.Exec(sqlStatement, id, file)
	if err != nil {
		panic(err)
	}

	return nil
}

func DeleteClips(projectID string) error {
	sqlStatement :=
		`DELETE
		 FROM clips
		 WHERE project_id=$1
		`
	_, err := db.Exec(sqlStatement, projectID)
	if err != nil {
		panic(err)
	}
	return nil
}
