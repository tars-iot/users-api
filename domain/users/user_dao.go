package users

import (
	_ "github.com/go-sql-driver/mysql" // mysql driver library

	usersdb "github.com/tars-iot/users-api/data-sources/mysql/users_db"
	dateutils "github.com/tars-iot/users-api/utils/date-utils"
	mysqlutils "github.com/tars-iot/users-api/utils/mysql_utils"

	"github.com/tars-iot/users-api/utils/errors"
)

const (
	queryInsertUser        = `INSERT INTO users (first_name, last_name, email, date_created, password, status) VALUES (?, ?, ?, ?, ?, ?)`
	queryGetUser           = `SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?`
	queryUpdateUser        = `UPDATE users SET first_name=? ,last_name=? ,email=? WHERE id=?`
	queryDeleteUser        = `DELETE FROM users WHERE id=?`
	queryFindUsersByStatus = `SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?`
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

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Password, user.Status)
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

// Update is function to update user to database
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

// Delete is function to Delete user from database
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

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	if err := usersdb.Client.Ping(); err != nil {
		panic(err)
	}
	stmt, err := usersdb.Client.Prepare(queryFindUsersByStatus)
	if err != nil {
		return nil, errors.InternalServerErr(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)

	if err != nil {
		return nil, errors.InternalServerErr(err.Error())
	}
	defer rows.Close()

	results := make([]User, 0)

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, mysqlutils.ParseError(err)
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NotFoundErr("No users found matching status")
	}
	return results, nil
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
