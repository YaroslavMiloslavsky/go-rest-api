package database

import (
	"github.com/YaroslavMiloslavsky/go-rest-api/api"
)

type DBInterface interface {
	GetAllUsers() (*[]api.UserFromDB, error)
}
