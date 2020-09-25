package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser is Handler to create new entry of user in database
func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement Me!!!")
}

// GetUser is Handler to perform get user details
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement Me!!!")
}
