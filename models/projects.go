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

func ReadProjectsforUser() (*[]dto.Project, error) {
	var p dto.Project
	var projects []dto.Project
	rows, err := db.Query(`SELECT id, name, song, created, status, url, openshot_id FROM projects`)
	if err != nil {
		fmt.Println("error querying")
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&p.ID, &p.Name, &p.Song, &p.Created, &p.Status, &p.OpenshotURL, &p.OpenshotID)
		if err != nil {
			// handle this error
			fmt.Println("error querying by row")
			return nil, err
		}
		projects = append(projects, p)
	}
	return &projects, nil
}
