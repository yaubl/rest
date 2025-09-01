package users

import (
	"bls/db"

	"github.com/julienschmidt/httprouter"
)

func RegisterRoutes(r *httprouter.Router, q *db.Queries) {
	service := NewService(q)
	controller := NewController(service)

	r.GET("/users", controller.GetAll)
	r.GET("/users/:id", controller.GetOne)
}
