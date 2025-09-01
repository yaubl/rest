package auth

import (
	"bls/api/middlewares"
	"bls/api/response"
	"bls/config"
	"net/http"
	"net/url"

	"github.com/julienschmidt/httprouter"
)

type Controller struct {
	service *Service
}

func NewController(s *Service) *Controller {
	return &Controller{service: s}
}

func (c *Controller) Me(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user, _ := middlewares.FromContext(r.Context())

	response.Json(w, http.StatusOK, user)
}

func (c *Controller) RedirectLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	params := url.Values{}
	params.Add("client_id", config.ClientID)
	params.Add("redirect_uri", config.RedirectURI)
	params.Add("response_type", "code")
	params.Add("scope", "identify")

	http.Redirect(w, r, "https://discord.com/api/oauth2/authorize?"+params.Encode(), http.StatusFound)
}

func (c *Controller) DiscordCallback(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	code := r.URL.Query().Get("code")
	if code == "" {
		response.Json(w, http.StatusBadRequest, response.Error{Error: "missing code"})
		return
	}

	user, sessionID, err := c.service.Callback(r.Context(), code)
	if err != nil {
		response.Json(w, http.StatusInternalServerError, response.Error{Error: err.Error()})
		return
	}

	response.Json(w, http.StatusOK, map[string]any{
		"user":  user,
		"token": sessionID,
	})
}
