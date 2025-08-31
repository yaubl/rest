package auth

import (
	"api/services"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	authRouter := router.Group("/auth")

	authRouter.GET("/login", services.AuthLogin)
	authRouter.GET("/callback", services.AuthCallback)
}
