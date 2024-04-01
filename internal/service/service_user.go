package service

import (
	"log"

	"github.com/YaroslavMiloslavsky/go-rest-api/api"
	"github.com/YaroslavMiloslavsky/go-rest-api/internal/database"
)

type UserServiceImpl struct {
	database *database.DBInterface
}

func CreateNewUserService() *UserServiceImpl {
	var service UserServiceImpl = UserServiceImpl{}

	var db database.DBInterface = database.CreateNewPostgresDB()
	service.database = &db

	return &service
}

func (u *UserServiceImpl) GetAll() (*api.UsersGetAll, error) {
	usersFromDB, err :=(*u.database).GetAllUsers()

	if err != nil {
		log.Printf("Could not get users: %v\n", err)
		return &api.UsersGetAll{}, nil
	}

	users := []api.UserDTO{}

	for _, user := range *usersFromDB {
		users = append(users, api.UserDTO{
			Username: user.Username,
			Age:      user.Age,
		})
	}

	response := api.UsersGetAll{}
	response.Users = users
	response.Count = len(*usersFromDB)

	return &response, nil
}
