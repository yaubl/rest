package auth

import (
	"api/db"
	"api/middlewares"
	"api/pkg/jwt"
	"api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	authRouter := router.Group("/auth")

	authRouter.GET("/login", Login)
	authRouter.GET("/callback", Callback)
}

func Login(c *gin.Context) {
	state := ""
	c.Redirect(http.StatusFound, services.GetDiscordOAuthURL(state))
}

func Callback(c *gin.Context) {
	ctx := middlewares.GetAppContext(c)
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing code"})
		return
	}

	user, err := services.Exchange(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := jwt.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	_, updateErr := ctx.DB.UpdateUser(ctx.Context, db.UpdateUserParams{ID: user.ID, Username: user.Username})
	if updateErr != nil {
		ctx.DB.CreateUser(ctx.Context, db.CreateUserParams{ID: user.ID, Username: user.Username})
	}

	c.JSON(http.StatusOK, token)
}
