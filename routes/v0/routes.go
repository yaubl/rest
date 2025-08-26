package v0

import "github.com/gin-gonic/gin"

func SetupRouter(router *gin.Engine) {
	v0Router := router.Group("/v0")

	// users
	v0Router.GET("users", getUsers)
	v0Router.POST("user", createUser) //note: this route is just to test
	// the user creation and it should NOT exist, remove ASAP OK?!?!?
}
