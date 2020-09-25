package users

import (
	"fmt"

	"github.com/tars-iot/users-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

// Save is the function to store data in database
func (user *User) Save() *errors.RestErr {
	result := userDB[user.ID]
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
	result := userDB[user.ID]
	if result == nil {
		return errors.NotFoundErr(fmt.Sprintf("User %s not found", user.ID))
	}

	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}
