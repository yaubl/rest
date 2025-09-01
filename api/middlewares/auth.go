package middlewares

import (
	"bls/api/response"
	"bls/config"
	"bls/db"
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/julienschmidt/httprouter"
)

type contextKey string

const userContextKey contextKey = "user"

func AuthMiddleware(q *db.Queries, next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			response.Json(w, http.StatusUnauthorized, response.Error{Error: "missing authorization header"})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Json(w, http.StatusUnauthorized, response.Error{Error: "malformed authorization header"})
			return
		}

		tokenStr := parts[1]

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return config.JwtSecret, nil
		})
		if err != nil || !token.Valid {
			response.Json(w, http.StatusUnauthorized, response.Error{Error: "invalid token"})
			return
		}

		claims, claimsOk1 := token.Claims.(jwt.MapClaims)
		userID, claimsOk2 := claims["user_id"].(string)
		if !claimsOk1 || !claimsOk2 {
			response.Json(w, http.StatusUnauthorized, response.Error{Error: "invalid token claims"})
			return
		}

		user, err := q.GetUser(r.Context(), userID)
		if err != nil {
			response.Json(w, http.StatusUnauthorized, response.Error{Error: "user not found"})
			return
		}

		ctx := context.WithValue(r.Context(), userContextKey, user)
		next(w, r.WithContext(ctx), ps)
	}
}

func FromContext(ctx context.Context) (db.User, bool) {
	user, ok := ctx.Value(userContextKey).(db.User)
	return user, ok
}
