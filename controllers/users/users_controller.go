package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tars-iot/users-api/domain/users"
	"github.com/tars-iot/users-api/services"
	"github.com/tars-iot/users-api/utils/errors"
)

// CreateUser is Handler to create new entry of user in database
func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		restErr := errors.BadRequestErr(err.Error())
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		fmt.Println(saveErr)
		c.JSON(saveErr.StatusCode, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

// GetUser is Handler to perform get user details
func GetUser(c *gin.Context) {

	userID, usrErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if usrErr != nil {
		err := errors.BadRequestErr("Invalid user ID")
		c.JSON(err.StatusCode, err)
		return
	}

	result, getErr := services.GetUser(userID)
	if getErr != nil {
		fmt.Println(getErr)
		c.JSON(getErr.StatusCode, getErr)
		return
	}
	c.JSON(http.StatusOK, result)
}
