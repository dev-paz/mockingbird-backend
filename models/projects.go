package models

import (
	"fmt"

	"github.com/mockingbird-backend/dto"

	_ "github.com/lib/pq"
)

func CreateProject(p *dto.Project) error {
	sqlStatement := `
	INSERT INTO projects (id, name, song, created, status, url, openshot_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id`
	id := ""
	err := db.QueryRow(sqlStatement, p.ID, p.Name, p.Song, p.Created, p.Status, p.OpenshotURL, p.OpenshotID).Scan(&id)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func ReadAllProjects(userID string) (*[]dto.Project, error) {
	var p dto.Project
	var projects []dto.Project
	sqlStatement := `
		WITH user_projects AS (
			SELECT p.* FROM clips AS c
			INNER JOIN projects AS p ON c.project_id = p.id
			WHERE c.user_id=$1
		)

		SELECT pv.id, pv.name, pv.status, pv.users, pv.clips, pv.song
			FROM user_projects as up
			INNER JOIN projects_view AS pv ON pv.id=up.id;
		`

	// u := dto.UserSlice{}

	rows, err := db.Query(sqlStatement, userID)
	if err != nil {
		fmt.Println("error querying")
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&p.ID, &p.Name, &p.Status, &p.Users, &p.Clips, &p.Song)
		if err != nil {
			// handle this error
			fmt.Println("error querying by row")
			return nil, err
		}
		fmt.Println(p)
		projects = append(projects, p)
	}
	return &projects, nil
}
