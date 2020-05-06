package chat

import (
	"net/http"

	"github.com/go-chi/chi"
)

type Router struct {
	Handler *Handler
	Hub 	*Hub
}

func NewRouter(handler *Handler, hub *Hub) *Router {
	return &Router{
		Handler: handler,
		Hub: hub,
	}
}

func (r *Router) Routes(c chi.Router) {
	c.Get("/", r.Handler.ListHandler)

	go r.Hub.Run()

	c.HandleFunc("/ws", func(w http.ResponseWriter, req *http.Request) {
		ServeWs(r.Hub, w, req)
	})
}
