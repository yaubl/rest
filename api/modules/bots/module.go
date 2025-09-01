package bots

import (
	"bls/api/middlewares"
	"bls/db"

	"github.com/julienschmidt/httprouter"
)

func RegisterRoutes(r *httprouter.Router, q *db.Queries) {
	service := NewService(q)
	controller := NewController(service)

	r.POST("/bots", middlewares.AuthMiddleware(q, controller.Create))
	r.GET("/bots", controller.GetAll)
	r.GET("/bots/:id", controller.GetOne)
}
