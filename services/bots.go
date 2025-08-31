package services

import (
	"api/db"
	"api/middlewares"
	"api/pkg/snowflake"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBot(c *gin.Context) {
	ctx := middlewares.GetAppContext(c)
	id := c.Param("id")

	bot, err := ctx.DB.GetBotByID(ctx.Context, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bot)
}

func CreateBot(c *gin.Context) {
	var input db.CreateBotParams
	user := c.GetString("user_id")
	ctx := middlewares.GetAppContext(c)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !snowflake.IsSnowflake(input.ID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bot id must be a valid snowflake"})
		return
	}

	input.Author = user

	bot, err := ctx.DB.CreateBot(ctx.Context, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, bot)
}
