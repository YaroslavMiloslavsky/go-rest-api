package handler

import (
	"encoding/json"
	"net/http"

	"github.com/YaroslavMiloslavsky/go-rest-api/api"
	"github.com/YaroslavMiloslavsky/go-rest-api/internal/database"
	header_const "github.com/YaroslavMiloslavsky/go-rest-api/pkg/headers"
)

var db database.DBInterface = database.CreateNewPostgresDB()


func GetAllUsers(r http.ResponseWriter, req *http.Request) {
	usersFromDB, err := db.GetAllUsers()
	if err != nil {
		log.Printf("Could not get users: %v\n", err)
		return
	}

	users := []api.UserDTO{}

	for _, user := range *usersFromDB {
		users = append(users, api.UserDTO{
			Username: user.Username,
			Age: user.Age,
		})
	}

	response := api.UsersGetAll{}
	response.Users = users
	response.Count = len(*usersFromDB)

	r.Header().Set(header_const.WebHeaderKeyContentType, header_const.WebHeaderValueContentTypeJson)
	encoder := json.NewEncoder(r)
	err = encoder.Encode(response)

	if err != nil {
		log.Printf("Error processing the response: %v\n", err)
		return
	}
}