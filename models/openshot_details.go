package models

import (
	"fmt"

	"github.com/mockingbird-backend/dto"
)

func ReadOpenshotDetails() (*dto.OpenshotDetails, error) {
	sqlStatement :=
		`SELECT  ip_address, username, password
		 FROM openshot_ip
		`
	var od dto.OpenshotDetails
	row := db.QueryRow(sqlStatement)
	err := row.Scan(&od.IP, &od.Username, &od.Password)
	if err != nil {
		fmt.Printf(err.Error())
		return nil, err
	}
	return &od, nil
}
