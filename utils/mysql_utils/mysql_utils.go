package mysql_utils

import (
	"github.com/go-sql-driver/mysql"
	"github.com/tars-iot/users-api/utils/errors"
	"strings"
)

const (
	errRowNotExist = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errRowNotExist) {
			return errors.NotFoundErr("User not found")
		}
		return errors.InternalServerErr("Error in parsing SQL response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.ConflictErr("User already exist")
	default:
		return errors.InternalServerErr("Error in parsing SQL response")
	}

}
