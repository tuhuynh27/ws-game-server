package chat

import (
	"net/http"

	"github.com/oddx-team/odd-game-server/pkg/json"
)

func (s *Services) ListHandler(w http.ResponseWriter, r *http.Request) {
	res := json.Response{ResponseWriter: w}

	chats, err := s.ListService()
	if err != nil {
		res.SendBadRequest(err.Error())
	}

	res.SendOK(chats)
}
