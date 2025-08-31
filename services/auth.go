package services

import (
	"api/config"
	"api/db"
	"api/middlewares"
	"api/pkg/jwt"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

type DiscordUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

func AuthLogin(c *gin.Context) {
	params := url.Values{}
	params.Add("client_id", config.ClientID)
	params.Add("redirect_uri", config.RedirectURI)
	params.Add("response_type", "code")
	params.Add("scope", "identify")

	c.Redirect(http.StatusFound, "https://discord.com/api/oauth2/authorize?"+params.Encode())
}

// note for myself: this works :D
func AuthCallback(c *gin.Context) {
	ctx := middlewares.GetAppContext(c)
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing code"})
		return
	}

	data := url.Values{}
	data.Set("client_id", config.ClientID)
	data.Set("client_secret", config.ClientSecret)
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", config.RedirectURI)

	req, _ := http.NewRequest("POST", "https://discord.com/api/oauth2/token", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != 200 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token exchange failed"})
		return
	}
	defer resp.Body.Close()

	var tokenResp struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "decode token failed"})
		return
	}

	reqUser, _ := http.NewRequest("GET", "https://discord.com/api/users/@me", nil)
	reqUser.Header.Set("Authorization", fmt.Sprintf("%s %s", tokenResp.TokenType, tokenResp.AccessToken))

	respUser, err := http.DefaultClient.Do(reqUser)
	if err != nil || respUser.StatusCode != 200 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "fetch user failed"})
		return
	}
	defer respUser.Body.Close()

	var user DiscordUser
	if err := json.NewDecoder(respUser.Body).Decode(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "decode user failed"})
		return
	}

	signed, err := jwt.GenerateJWT(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to sign JWT"})
		return
	}

	dbUser, err := ctx.DB.CreateUser(ctx.Context, db.CreateUserParams{ID: user.ID, Username: user.Username})
	if err != nil {
		dbUser, _ = ctx.DB.UpdateUsername(ctx.Context, db.UpdateUsernameParams{ID: user.ID, Username: user.Username})
	}

	c.JSON(http.StatusOK, gin.H{
		"token": signed,
		"user":  dbUser,
	})
}
