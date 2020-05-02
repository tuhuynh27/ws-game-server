package chat

import (
	"github.com/oddx-team/odd-game-server/pkg/json"
	"net/http"
)

func ListHandler(w http.ResponseWriter, r *http.Request) {
	res := json.Response{ResponseWriter: w}

	chats, err := ListService()
	if err != nil {
		res.SendBadRequest(err.Error())
	}

	res.SendOK(chats)
}
