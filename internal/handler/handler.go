package handler

import (
	"net/http"

	"github.com/YaroslavMiloslavsky/go-rest-api/internal/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var (
	log = logger.CreateNew()
)

func Handler(r *http.Handler) {
	muxRouter, ok := (*r).(*chi.Mux)

	if !ok {
		// If provided type is not of handler type, crash the service
		log.Fatal("Router does not implement *http.Handler interface")
	}

	muxRouter.Use(middleware.StripSlashes)
	muxRouter.Route("/users", func(r chi.Router) {
		log.Println("Mapping routes to handler")
		// TODO: add Auth -> only ADMIN can execute this
		r.Get("/all", GetAllUsers)
	});
	log.Println("Routes were mapped")
}
