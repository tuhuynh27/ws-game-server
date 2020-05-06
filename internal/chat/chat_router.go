package chat

import (
	"net/http"

	"github.com/go-chi/chi"
)

type Router struct {
	Handler *Handler
	Hub 	*Hub
	Routes 	http.Handler
}

func NewRouter(handler *Handler, hub *Hub) *Router {
	return &Router{
		Handler: handler,
		Hub: hub,
		Routes: getRoutes(handler, hub),
	}
}

func getRoutes(handler *Handler, hub *Hub) http.Handler {
	c := chi.NewRouter()
	c.Get("/", handler.ListHandler)
	c.HandleFunc("/ws", func(w http.ResponseWriter, req *http.Request) {
		ServeWs(hub, w, req)
	})

	return c
}
