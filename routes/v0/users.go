package v0

import (
	"api/services"

	"github.com/gin-gonic/gin"
)

func RegisterUsersRoutes(router *gin.RouterGroup) {
	users := router.Group("/users")

	users.GET("/:id", services.GetUser)
	users.GET("/", services.GetUsers)
}
