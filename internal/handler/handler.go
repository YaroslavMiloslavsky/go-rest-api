package handler

import (
	"net/http"

	"github.com/YaroslavMiloslavsky/go-rest-api/internal/logger"
)

var (
	log         = logger.CreateNew()
	userHandler = CreateNewUserHandler()
)

func Handler(r *http.Handler) {
	var router *http.ServeMux
	router, ok := (*r).(*http.ServeMux)

	if !ok {
		log.Fatal("Router does not implement *http.Handler interface")
	}

	router.HandleFunc("GET /users/all", userHandler.GetAllUsers)
	log.Println("Routes were mapped")
}
