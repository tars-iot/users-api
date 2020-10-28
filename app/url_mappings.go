package app

import (
	"github.com/tars-iot/users-api/controllers/ping"
	"github.com/tars-iot/users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("tars-iot/v1/users/:user_id", users.Get)

	router.POST("tars-iot/v1/users", users.Create)

	router.PUT("tars-iot/v1/users/:user_id", users.Update)

	router.PATCH("tars-iot/v1/users/:user_id", users.Update)

	router.DELETE("tars-iot/v1/users/:user_id", users.Delete)
}
