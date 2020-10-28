package users

import (
	_ "github.com/go-sql-driver/mysql"

	usersdb "github.com/tars-iot/users-api/data-sources/mysql/users_db"
	dateutils "github.com/tars-iot/users-api/utils/date-utils"
	mysqlutils "github.com/tars-iot/users-api/utils/mysql_utils"

	"github.com/tars-iot/users-api/utils/errors"
)

const (
	queryInsertUser = `INSERT INTO users (first_name, last_name, email, date_created) VALUES (?, ?, ?, ?)`
	queryGetUser    = `SELECT id, first_name, last_name, email, date_created from users WHERE id=?`
	queryUpdateUser = `UPDATE users SET first_name=? ,last_name=? ,email=? WHERE id=?`
	queryDeleteUser = `DELETE FROM users WHERE id=?`
)

// Save is the function to store data in database
func (user *User) Save() *errors.RestErr {
	if err := usersdb.Client.Ping(); err != nil {
		panic(err)
	}
	user.DateCreated = dateutils.GetNowString()
	stmt, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.InternalServerErr(err.Error())
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		return mysqlutils.ParseError(saveErr)
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.InternalServerErr(err.Error())
	}
	user.ID = userID

	return nil
}

// Get function is used to fetch data from database respect to ID
func (user *User) Get() *errors.RestErr {

	if err := usersdb.Client.Ping(); err != nil {
		panic(err)
	}
	stmt, err := usersdb.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.InternalServerErr(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		return mysqlutils.ParseError(getErr)
	}
	return nil
}

func (user *User) Update() *errors.RestErr {
	if err := usersdb.Client.Ping(); err != nil {
		panic(err)
	}
	stmt, err := usersdb.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.InternalServerErr(err.Error())
	}
	defer stmt.Close()
	_, updateErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if updateErr != nil {
		return mysqlutils.ParseError(updateErr)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	if err := usersdb.Client.Ping(); err != nil {
		panic(err)
	}
	stmt, err := usersdb.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.InternalServerErr(err.Error())
	}
	defer stmt.Close()
	_, deleteErr := stmt.Exec(user.ID)
	if deleteErr != nil {
		return mysqlutils.ParseError(deleteErr)
	}
	return nil
}

//SQL QUERY:
//	CREATE TABLE:
// CREATE TABLE `users` (
// `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
// `first_name` VARCHAR(20) DEFAULT NULL,
// `last_name` VARCHAR(20) DEFAULT NULL,
// `email` VARCHAR(40),
// `date_created` VARCHAR(20) DEFAULT NULL,
// UNIQUE KEY `email_UNIQUE` (`email`) USING HASH,
// PRIMARY KEY (`id`)
// );
