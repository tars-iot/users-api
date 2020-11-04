package services

import (
	"github.com/tars-iot/users-api/domain/users"
	crypto_utils "github.com/tars-iot/users-api/utils/crypto-utils"
	"github.com/tars-iot/users-api/utils/errors"
)

// CreateUser is service funntion to store the user data
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.Password = crypto_utils.GetMD5(user.Password)
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

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}

	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func DeleteUser(userID int64) *errors.RestErr {
	current, err := GetUser(userID)
	if err != nil {
		return err
	}

	if err := current.Delete(); err != nil {
		return err
	}
	return nil
}

func Search(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
