package v0

import (
	"api/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserInput struct {
	Username string `json:"username" binding:"required"`
}

func getUsers(c *gin.Context) {
	ctx := middlewares.GetAppContext(c)

	users, err := ctx.DB.ListUsers(ctx.Context)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func createUser(c *gin.Context) {
	var input CreateUserInput
	ctx := middlewares.GetAppContext(c)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ctx.DB.CreateUser(ctx.Context, input.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
