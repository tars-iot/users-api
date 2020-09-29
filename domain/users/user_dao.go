package users

import (
	"fmt"

	usersdb "github.com/tars-iot/users-api/data-sources/postgres/users-db"

	dateutils "github.com/tars-iot/users-api/utils/date-utils"
	"github.com/tars-iot/users-api/utils/errors"
)

const (
	queryInsertUser = `INSERT INTO users (first_name, last_name, email, date_created) VALUES (?, ?, ?, ?)`
)

var (
	userDB = make(map[int64]*User)
)

// Save is the function to store data in database
func (user *User) Save() *errors.RestErr {
	result := userDB[user.ID]
	user.DateCreated = dateutils.GetNowString()
	stmt, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.InternalServerErr(err.Error())
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		return errors.InternalServerErr(err.Error())
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.InternalServerErr(err.Error())
	}

	user.ID = userID
	if result != nil {
		if result.Email == user.Email {
			return errors.ConflictErr(fmt.Sprintf("Email: %s already exist", user.Email))
		}
		return errors.ConflictErr(fmt.Sprintf("User: %d already exist", user.ID))
	}

	userDB[user.ID] = user
	return nil
}

// Get function is used to fetch data from database respect to ID
func (user *User) Get() *errors.RestErr {

	if err := usersdb.Client.Ping(); err != nil {
		panic(err)
	}

	result := userDB[user.ID]
	if result == nil {
		return errors.NotFoundErr(fmt.Sprintf("User %d not found", user.ID))
	}

	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}
