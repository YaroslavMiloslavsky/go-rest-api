package main

import (
	"net/http"

	"github.com/YaroslavMiloslavsky/go-rest-api/internal/handler"
	"github.com/YaroslavMiloslavsky/go-rest-api/internal/logger"
	"github.com/YaroslavMiloslavsky/go-rest-api/pkg/utils"
	"github.com/YaroslavMiloslavsky/go-rest-api/internal/middleware"
)


func main() {
	var router http.Handler
	logger := logger.CreateNew()

	router = http.NewServeMux()
	handler.Handler(&router)
	mStack := middleware.CreateMiddlewareStack(
		middleware.StripSlashes,
	)

	serverURL := utils.GetServerURL()
	logger.Println("about to serve " + serverURL)

	http.ListenAndServe(serverURL, 	mStack(router))
}
