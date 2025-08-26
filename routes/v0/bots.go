package v0

import (
	"api/middlewares"
	"api/services"

	"github.com/gin-gonic/gin"
)

func RegisterBotsRoutes(router *gin.RouterGroup) {
	bots := router.Group("/bots")

	bots.GET("/:id", services.GetBot)
	bots.POST("/", middlewares.AuthMiddleware(), services.CreateBot)
}
