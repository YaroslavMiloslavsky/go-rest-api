package handler

import (
	"encoding/json"
	"net/http"

	"github.com/YaroslavMiloslavsky/go-rest-api/internal/service"
	"github.com/YaroslavMiloslavsky/go-rest-api/pkg/header_const"
)

var userService service.UserServiceInterface = service.CreateNewUserService()

func GetAllUsers(r http.ResponseWriter, req *http.Request) {
	response, err := userService.GetAll()

	if err != nil {
		log.Printf("Could not fetch users from service: %s\n", err)
	}

	r.Header().Set(header_const.WebHeaderKeyContentType, header_const.WebHeaderValueContentTypeJson)
	encoder := json.NewEncoder(r)
	err = encoder.Encode(response)

	if err != nil {
		log.Printf("Error processing the response: %v\n", err)
		return
	}
}