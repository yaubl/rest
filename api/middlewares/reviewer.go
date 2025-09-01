package middlewares

import (
	"bls/api/response"
	"bls/config"
	"bls/db"
	"net/http"
	"slices"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func ReviewerMiddleware(q *db.Queries, next httprouter.Handle) httprouter.Handle {
	return AuthMiddleware(q, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		user, ok := FromContext(r.Context())
		if !ok {
			response.Json(w, http.StatusUnauthorized, response.Error{Error: "unauthorized"})
			return
		}

		reviewers := strings.Split(config.Reviewers, ",")

		if !slices.Contains(reviewers, user.ID) {
			response.Json(w, http.StatusForbidden, response.Error{Error: "forbidden"})
			return
		}

		next(w, r, ps)
	})
}
