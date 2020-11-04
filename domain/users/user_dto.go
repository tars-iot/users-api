package users

import (
	"strings"

	"github.com/tars-iot/users-api/utils/errors"
)

const (
	StatusActive = "active"
)

type Users []User

// User is a structure of user data model
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

// Validate is used to validate manadated data in request
func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.BadRequestErr("Missing manadatory parameter: email")
	}
	if user.FirstName == "" {
		return errors.BadRequestErr("Missing manadatory parameter: first_name")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.BadRequestErr("Invalid password")
	}
	return nil
}
