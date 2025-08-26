package v0

import "github.com/gin-gonic/gin"

func Register(router *gin.Engine) {
	v0Router := router.Group("/v0")

	// users
	RegisterUsersRoutes(v0Router)
	// bots
	RegisterBotsRoutes(v0Router)
}
