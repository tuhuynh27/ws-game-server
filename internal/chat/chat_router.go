package chat

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (s *Services) Routes(r chi.Router) {
	r.Get("/", s.ListHandler)

	go s.Run()

	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ServeWs(s.Hub, w, r)
	})
}
