package usersdb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // github.com/lib/pq is postgres package
)

// const (
// 	postgresUsersdbUsername = "postgresUsersdbUsername"
// 	postgresUsersdbPassword = "postgresUsersdbPassword"
// 	postgresUsersdbHost     = "postgresUsersdbHost"
// 	postgresUsersdbSchema   = "postgresUsersdbSchema"
// )

//Client is sql sql connection string
var (
	Client *sql.DB

	// username = os.Getenv("postgresUsersdbUsername")
	// password = os.Getenv("postgresUsersdbPassword")
	// host     = os.Getenv("postgresUsersdbHost")
	// schema   = os.Getenv("postgresUsersdbSchema")
	username = "tarsiot"
	password = "tarsiot"
	host     = "localhost"
	schema   = "users_db"
)

func init() {
	var err error
	dataSourceConnectionString := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		username, password, host, schema,
	)
	Client, err = sql.Open("postgres", dataSourceConnectionString)
	if err != nil {
		panic(err)
	}

	// Set the maximum number of concurrently open connections (in-use + idle)
	// to 5. Setting this to less than or equal to 0 will mean there is no
	// maximum limit (which is also the default setting).
	Client.SetMaxOpenConns(5)

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("Configured database successfully!!!")
}
