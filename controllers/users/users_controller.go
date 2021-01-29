package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/halimakibb/bookstore_users-api/domain/users"
	"github.com/halimakibb/bookstore_users-api/services"
	"github.com/halimakibb/bookstore_users-api/utils/errors"
)

// GetUser from service
func GetUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userID)
	if getErr != nil {
		err := errors.NewBadRequestError("user id not found")
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

// CreateUser API
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// SearchUser for filtering
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
