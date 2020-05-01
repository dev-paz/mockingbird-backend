package models

import (
	"fmt"

	"github.com/mockingbird-backend/dto"

	_ "github.com/lib/pq"
)

func CreateProject(p *dto.ProjectDB) error {
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

func UpdateProjectStatus(id string, status string, exportID string) error {
	sqlStatement := `
	UPDATE projects
	SET status = $2, export_id = $3
	WHERE id=$1`
	_, err := db.Exec(sqlStatement, id, status, exportID)
	if err != nil {
		panic(err)
	}
	return nil
}

func ReadProject(projectID string) (*dto.Project, error) {
	sqlStatement :=
		`SELECT  id, name, status, clips, song, openshot_id
		 FROM projects_view
		 WHERE id=$1
		`
	var p dto.Project
	row := db.QueryRow(sqlStatement, projectID)
	err := row.Scan(&p.ID, &p.Name, &p.Status, &p.Clips, &p.Song, &p.OpenshotID)
	if err != nil {
		fmt.Printf(err.Error())
		return nil, err
	}
	return &p, nil
}

func DeleteProject(projectID string) error {
	sqlStatement :=
		`DELETE
		 FROM projects
		 WHERE id=$1
		`
	_, err := db.Exec(sqlStatement, projectID)
	if err != nil {
		panic(err)
	}
	return nil
}

func ReadAllProjects(userID string) (*[]dto.Project, error) {
	var p dto.Project
	var projects []dto.Project
	sqlStatement := `
		WITH user_projects AS (
			SELECT distinct p.* FROM clips AS c
			INNER JOIN projects AS p ON c.project_id = p.id
			WHERE c.user_id=$1
		)

		SELECT pv.id, pv.name, pv.status, pv.clips, pv.song, pv.openshot_id
			FROM user_projects as up
			INNER JOIN projects_view AS pv ON pv.id=up.id;
		`
	rows, err := db.Query(sqlStatement, userID)
	if err != nil {
		fmt.Println("error querying")
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&p.ID, &p.Name, &p.Status, &p.Clips, &p.Song, &p.OpenshotID)
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
