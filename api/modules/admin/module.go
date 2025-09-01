package admin

import (
	"bls/api/middlewares"
	"bls/db"

	"github.com/julienschmidt/httprouter"
)

func RegisterRoutes(r *httprouter.Router, q *db.Queries) {
	service := NewService(q)
	controller := NewController(service)

	// bots
	r.PATCH("/admin/bots/:id/approve", middlewares.ReviewerMiddleware(q, controller.ApproveBot))
	r.PATCH("/admin/bots/:id/queue", middlewares.ReviewerMiddleware(q, controller.QueueBot))
	r.PATCH("/admin/bots/:id/deny", middlewares.ReviewerMiddleware(q, controller.DenyBot))

	// users
	// todo: users
}
