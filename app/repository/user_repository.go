package repository

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"bitbucket.org/adhikag24/golang-api/app/domain"
	"bitbucket.org/adhikag24/golang-api/app/utils"
)

type Repository struct {
}

func (repo *Repository) GetAllUsers() ([]*domain.User, *utils.ApplicationError) {
	db := utils.CreateConnection()

	defer db.Close()

	var users []*domain.User

	sqlStatement := `SELECT * FROM users`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var user = &domain.User{}
		// unmarshal the row object to user
		err = rows.Scan(&user.UserID, &user.Name)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		// append the user in the users slice
		users = append(users, user)

	}

	if err == nil {
		return users, nil
	} else {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("errror: %v", err),
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}

}

func (repo *Repository) GetUser(userID string) (*domain.User, *utils.ApplicationError) {

	db := utils.CreateConnection()
	defer db.Close()

	var user = &domain.User{}

	sqlStatement := `SELECT * FROM users WHERE userid=$1`
	row := db.QueryRow(sqlStatement, userID)
	err := row.Scan(&user.UserID, &user.Name)

	switch err {
	case sql.ErrNoRows:
		fmt.Printf("Error in Get Data: %v", err)
		return nil, &utils.ApplicationError{
			Message:    "userid not found",
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	case nil:
		return user, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	if err == nil {
		return user, nil
	} else {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("error: %v", err),
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}
}
