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

func getUserID(userIDParam string) (int64, *errors.RestErr) {
	userID, usrErr := strconv.ParseInt(userIDParam, 10, 64)

	if usrErr != nil {
		return 0, errors.NotFoundErr("Invalid user ID")
	}
	return userID, nil

}

// Create is Handler to create new entry of user in database
func Create(c *gin.Context) {
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
	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

// Get is Handler to perform get user details
func Get(c *gin.Context) {

	userID, idErr := getUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.StatusCode, idErr)
		return
	}

	result, getErr := services.GetUser(userID)
	if getErr != nil {
		fmt.Println(getErr)
		c.JSON(getErr.StatusCode, getErr)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func Update(c *gin.Context) {
	userID, idErr := getUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.StatusCode, idErr)
		return
	}

	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		restErr := errors.BadRequestErr(err.Error())
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	user.ID = userID

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func Delete(c *gin.Context) {
	userID, idErr := getUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.StatusCode, idErr)
		return
	}
	err := services.DeleteUser(userID)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.Status(http.StatusNoContent)
}

// Search is function used to query based on probvided queryu parameter
func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := services.Search(status)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))

}
