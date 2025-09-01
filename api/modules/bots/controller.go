package bots

import (
	"bls/api/middlewares"
	"bls/api/response"
	"bls/db"
	"bls/pkg/snowflake"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type Controller struct {
	service *Service
}

func NewController(s *Service) *Controller {
	return &Controller{s}
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var input db.CreateBotParams

	user, _ := middlewares.FromContext(r.Context())
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		response.Json(w, http.StatusBadRequest, response.Error{Error: err.Error()})
		return
	}

	input.Author = user.ID
	bot, err := c.service.Create(r.Context(), input)
	if err != nil {
		response.Json(w, http.StatusInternalServerError, response.Error{Error: err.Error()})
		return
	}

	response.Json(w, http.StatusCreated, bot)
}

func (c *Controller) GetOne(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !snowflake.IsSnowflake(id) {
		response.Json(w, http.StatusBadRequest, response.Error{Error: "invalid snowflake"})
		return
	}

	bot, err := c.service.GetOne(r.Context(), id)
	if err != nil {
		response.Json(w, http.StatusNotFound, response.Error{Error: err.Error()})
		return
	}

	response.Json(w, http.StatusOK, bot)
}

func (c *Controller) GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	q := r.URL.Query()

	limitStr := q.Get("limit")
	offsetStr := q.Get("offset")

	if limitStr == "" {
		limitStr = "15"
	}
	if offsetStr == "" {
		offsetStr = "0"
	}

	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		limit = 15
	}

	offset, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		offset = 0
	}

	bots, err := c.service.GetAll(r.Context(), limit, offset)
	if err != nil {
		response.Json(w, http.StatusInternalServerError, response.Error{Error: err.Error()})
		return
	}

	response.Json(w, http.StatusOK, bots)
}
