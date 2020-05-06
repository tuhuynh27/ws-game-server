package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	"github.com/oddx-team/odd-game-server/config"
	"github.com/oddx-team/odd-game-server/internal/chat"
)

func main() {
	cfg := config.Load()
	mongoConn := cfg.NewMongoConnection(cfg.Mongo.Host, cfg.Mongo.DatabaseName)

	chatService := chat.NewService(mongoConn)
	chatHandler := chat.NewHandler(chatService)
	chatWsHub := chat.NewHub(chatService)
	chatRouter := chat.NewRouter(chatHandler, chatWsHub)

	r := serveHTTP()
	go chatRouter.Hub.Run()
	r.Mount("/api/v1/chat", chatRouter.Routes)

	log.Println("Started at port " + cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}

func serveHTTP() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	corsOptions := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(corsOptions.Handler)
	r.Use(middleware.Timeout(30 * time.Second))
	return r
}
