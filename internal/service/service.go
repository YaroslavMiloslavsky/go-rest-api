package service

import "github.com/YaroslavMiloslavsky/go-rest-api/api"

type UserServiceInterface interface {
	GetAll() (*api.UsersGetAll, error)
}