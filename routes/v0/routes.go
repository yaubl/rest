package v0

import "github.com/gin-gonic/gin"

func SetupRouter(router *gin.Engine) {
	v0Router := router.Group("/v0")

	v0Router.GET("users", getUsers)
}
