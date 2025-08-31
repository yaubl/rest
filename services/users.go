package services

import (
	"api/db"
	"api/middlewares"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	ctx := middlewares.GetAppContext(c)
	id := c.Param("id")

	user, err := ctx.DB.GetUser(ctx.Context, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUsers(c *gin.Context) {
	ctx := middlewares.GetAppContext(c)
	limitStr, limitOk := c.GetQuery("limit")
	offsetStr, offsetOk := c.GetQuery("offset")

	if !limitOk || limitStr == "" {
		limitStr = "15"
	}
	if !offsetOk || offsetStr == "" {
		offsetStr = "0"
	}

	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		limit = 15
	}

	offset, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		offset = 0
	}

	users, err := ctx.DB.ListUsers(ctx.Context, db.ListUsersParams{
		Limit:  limit,
		Offset: offset,
	})

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
