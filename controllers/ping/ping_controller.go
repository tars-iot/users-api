package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping is handler which reply pong as response for /ping url
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "Pong")
}
