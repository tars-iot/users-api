package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication is used to initialise all views used by external endpoints
func StartApplication() {
	mapUrls()
	router.Run(":8080")
}
