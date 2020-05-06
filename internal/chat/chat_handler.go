package chat

import (
	"net/http"

	"github.com/oddx-team/odd-game-server/pkg/json"
)

type Handler struct {
	Service IService
}

func NewHandler(service IService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) ListHandler(w http.ResponseWriter, _ *http.Request) {
	res := json.Response{ResponseWriter: w}

	chats, err := h.Service.ListService()
	if err != nil {
		res.SendBadRequest(err.Error())
	}

	res.SendOK(chats)
}
