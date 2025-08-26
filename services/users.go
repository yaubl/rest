package services

import (
	"api/db"
	"api/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	ctx := middlewares.GetAppContext(c)
	id := c.Param("id")

	user, err := ctx.DB.GetUserByID(ctx.Context, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUsers(c *gin.Context) {
	ctx := middlewares.GetAppContext(c)

	users, err := ctx.DB.ListUsers(ctx.Context)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var input db.CreateUserParams
	ctx := middlewares.GetAppContext(c)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ctx.DB.CreateUser(ctx.Context, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
