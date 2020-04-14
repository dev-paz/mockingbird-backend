package models

import (
	"fmt"

	"github.com/mockingbird-backend/dto"
)

func CreateUser(u *dto.User) error {
	sqlStatement := `
	INSERT INTO users (id, firebase_user_id, profile_photo_url, email, name)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id`
	id := ""
	err := db.QueryRow(sqlStatement, u.ID, u.FirebaseUserID, u.ProfilePhotoURL, u.Email, u.Name).Scan(&id)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func ReadUser(firebaseUserId string) (*dto.User, error) {
	sqlStatement := `SELECT id, firebase_user_id, profile_photo_url, email, name FROM users WHERE firebase_user_id=$1;`
	var u dto.User
	row := db.QueryRow(sqlStatement, firebaseUserId)
	err := row.Scan(&u.ID, &u.FirebaseUserID, &u.ProfilePhotoURL, &u.Email, &u.Name)
	if err != nil {
		fmt.Printf(err.Error())
		return nil, err
	}
	return &u, nil
}
