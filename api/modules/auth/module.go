package auth

import (
	"bls/api/middlewares"
	"bls/db"

	"github.com/julienschmidt/httprouter"
)

func RegisterRoutes(r *httprouter.Router, q *db.Queries) {
	service := NewService(q)
	controller := NewController(service)

	r.GET("/auth/me", middlewares.AuthMiddleware(q, controller.Me))
	r.GET("/auth/login", controller.RedirectLogin)
	r.GET("/auth/callback", controller.DiscordCallback)
}
