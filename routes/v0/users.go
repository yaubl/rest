package v0

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, "mocked users!")
}
