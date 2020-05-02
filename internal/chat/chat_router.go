package chat

import (
	"github.com/go-chi/chi"
	"net/http"
)

func Routes(r chi.Router) {
	r.Get("/", ListHandler)

	hub := NewHub()
	go hub.Run()

	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ServeWs(hub, w, r)
	})
}
