package api

import (
	"bls/api/modules/auth"
	"bls/api/modules/bots"
	"bls/api/modules/users"

	"github.com/julienschmidt/httprouter"
)

func NewRouter() *httprouter.Router {
	router := httprouter.New()

	bots.RegisterRoutes(router, queries)
	users.RegisterRoutes(router, queries)
	auth.RegisterRoutes(router, queries)

	return router
}
