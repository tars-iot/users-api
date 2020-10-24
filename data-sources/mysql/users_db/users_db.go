package usersdb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//TODO: Add to OS.Env group for production
const (
	mysqlUsersUsername = "root"
	mysqlUsersPassword = "tars-iot-database"
	mysqlUsersHost     = "192.168.49.2:30423"
	mysqlUsersSchema   = "users_db"
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
