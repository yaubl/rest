package auth

import (
	"bls/config"
	"bls/db"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Service struct {
	q *db.Queries
}

func NewService(q *db.Queries) *Service {
	return &Service{q}
}

type discordTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type discordUserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (s *Service) Me(ctx context.Context, id string) (db.User, error) {
	user, err := s.q.GetUser(ctx, id)
	if err != nil {
		return db.User{}, errors.New("user not found")
	}

	return user, nil
}

func (s *Service) Callback(ctx context.Context, code string) (db.User, string, error) {
	if code == "" {
		return db.User{}, "", errors.New("empty code")
	}

	form := url.Values{}
	form.Set("client_id", config.ClientID)
	form.Set("client_secret", config.ClientSecret)
	form.Set("grant_type", "authorization_code")
	form.Set("code", code)
	form.Set("redirect_uri", config.RedirectURI)

	req, _ := http.NewRequestWithContext(ctx, "POST", "https://discord.com/api/oauth2/token", bytes.NewBufferString(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return db.User{}, "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return db.User{}, "", errors.New("failed to get discord token")
	}

	var tokenRes discordTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenRes); err != nil {
		return db.User{}, "", err
	}

	userReq, _ := http.NewRequestWithContext(ctx, "GET", "https://discord.com/api/users/@me", nil)
	userReq.Header.Set("Authorization", "Bearer "+tokenRes.AccessToken)

	userResp, err := http.DefaultClient.Do(userReq)
	if err != nil {
		return db.User{}, "", err
	}
	defer userResp.Body.Close()

	if userResp.StatusCode != http.StatusOK {
		return db.User{}, "", errors.New("failed to fetch discord user")
	}

	var discordUser discordUserResponse
	if err := json.NewDecoder(userResp.Body).Decode(&discordUser); err != nil {
		return db.User{}, "", err
	}

	user, err := s.q.CreateUser(ctx, db.CreateUserParams{
		ID:       discordUser.ID,
		Username: discordUser.Username,
	})
	if err != nil {
		user, err = s.q.UpdateUsername(ctx, db.UpdateUsernameParams{ID: discordUser.ID, Username: discordUser.Username})
		if err != nil {
			return db.User{}, "", err
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})
	sessionID, err := token.SignedString(config.JwtSecret)
	if err != nil {
		return db.User{}, "", err
	}

	return user, sessionID, nil
}
