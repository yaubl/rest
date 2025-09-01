package admin

import (
	"bls/api/middlewares"
	"bls/api/response"
	"bls/logger"
	"bls/pkg/snowflake"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Controller struct {
	service *Service
}

func NewController(s *Service) *Controller {
	return &Controller{s}
}

func (c *Controller) ApproveBot(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user, _ := middlewares.FromContext(r.Context())
	id := p.ByName("id")

	if !snowflake.IsSnowflake(id) {
		response.Json(w, http.StatusBadRequest, response.Error{Error: "invalid snowflake"})
		return
	}

	bot, err := c.service.UpdateBotStatus(r.Context(), id, "approved")
	if err != nil {
		response.Json(w, http.StatusBadRequest, response.Error{Error: "invalid snowflake"})
		return
	}

	logger.Log.Infof("approved bot %s (%s). issued by %s (%s)", bot.Name, bot.ID, user.Username, user.ID)

	response.Json(w, http.StatusOK, bot)
}

func (c *Controller) QueueBot(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user, _ := middlewares.FromContext(r.Context())
	id := p.ByName("id")

	if !snowflake.IsSnowflake(id) {
		response.Json(w, http.StatusBadRequest, response.Error{Error: "invalid snowflake"})
		return
	}

	bot, err := c.service.UpdateBotStatus(r.Context(), id, "pending")
	if err != nil {
		response.Json(w, http.StatusBadRequest, response.Error{Error: "invalid snowflake"})
		return
	}

	logger.Log.Infof("queued bot %s (%s). issued by %s (%s)", bot.Name, bot.ID, user.Username, user.ID)

	response.Json(w, http.StatusOK, bot)
}

func (c *Controller) DenyBot(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user, _ := middlewares.FromContext(r.Context())
	id := p.ByName("id")

	if !snowflake.IsSnowflake(id) {
		response.Json(w, http.StatusBadRequest, response.Error{Error: "invalid snowflake"})
		return
	}

	bot, err := c.service.UpdateBotStatus(r.Context(), id, "denied")
	if err != nil {
		response.Json(w, http.StatusBadRequest, response.Error{Error: "invalid snowflake"})
		return
	}

	logger.Log.Infof("denied bot %s (%s). issued by %s (%s)", bot.Name, bot.ID, user.Username, user.ID)

	response.Json(w, http.StatusOK, bot)
}
