package main

import (
	"net/http"

	"github.com/YaroslavMiloslavsky/go-rest-api/internal/handler"
	"github.com/YaroslavMiloslavsky/go-rest-api/internal/logger"
	"github.com/YaroslavMiloslavsky/go-rest-api/pkg/utils"
	"github.com/go-chi/chi/v5"
)

var r http.Handler

func main() {
	logger := logger.CreateNew()

	r = chi.NewRouter()
	handler.Handler(&r)
	serverURL := utils.GetServerURL()
	logger.Println("about to serve " + serverURL)

	http.ListenAndServe(serverURL, r)
}
