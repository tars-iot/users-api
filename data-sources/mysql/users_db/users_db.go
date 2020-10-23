package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//TODO: Add to OS.Env group for production
const (
	mysqlUsersUsername = "root"
	mysqlUsersPassword = "my-super-secret-password"
	mysqlUsersHost     = "10.32.131.35:31136"
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
