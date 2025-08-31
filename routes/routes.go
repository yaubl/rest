package routes

import (
	"api/routes/auth"
	"api/routes/v0"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// v0 routes
	v0.Register(router)
	// auth routes
	auth.Register(router)
}
