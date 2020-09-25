package app

import (
	"github.com/tars-iot/users-api/controllers/ping"
	"github.com/tars-iot/users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("tars-iot/v1/users/:users_id", users.GetUser)

	router.POST("tars-iot/v1/users", users.CreateUser)
}