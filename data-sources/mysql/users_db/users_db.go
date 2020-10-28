package usersdb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // mysql driver
)

//TODO: Add to OS.Env group for production
const (
	mysqlUsersUsername = "uys0w7fpuqdx0u3d"
	mysqlUsersPassword = "NOu08xSJmJJRnvXkQvzP"
	mysqlUsersHost     = "bvfmbfrbdycoxsx0ibvv-mysql.services.clever-cloud.com"
	mysqlUsersSchema   = "bvfmbfrbdycoxsx0ibvv"
)

var (
	Client *sql.DB
)

func init() {
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		mysqlUsersUsername,
		mysqlUsersPassword,
		mysqlUsersHost,
		mysqlUsersSchema,
	)
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database Successfully configured")
}
