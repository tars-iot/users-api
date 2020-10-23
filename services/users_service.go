package services

import (
	"github.com/tars-iot/users-api/domain/users"
	"github.com/tars-iot/users-api/utils/errors"
)

// CreateUser is service funntion to store the user data
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUser is service funntion to store the user data
func GetUser(userID int64) (*users.User, *errors.RestErr) {

	user := users.User{ID: userID}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	current.FirstName = user.FirstName
	current.LastName = user.LastName
	current.Email = user.Email

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}
