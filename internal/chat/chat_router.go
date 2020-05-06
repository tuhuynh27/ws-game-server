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
		Routes: getRoutes(),
	}
}

func getRoutes() http.Handler {
	c := chi.NewRouter()
	c.Get("/", r.Handler.ListHandler)
	c.HandleFunc("/ws", func(w http.ResponseWriter, req *http.Request) {
		ServeWs(r.Hub, w, req)
	})

	return c
}
