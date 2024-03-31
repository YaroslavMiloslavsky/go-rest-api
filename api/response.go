package api

import "time"

// Get user as appears in DB
type UserFromDB struct {
	Id       int       `json:"user_id"`
	Username string    `json:"username"`
	Age      int       `json:"age"`
	Created  time.Time `json:"created_at"`
	Updated  time.Time `json:"updated_at"`
}

// User without PII
type UserDTO struct {
	Username string `json:"username"`
	Age      int    `json:"age"`
}

// Get all users DTO
type UsersGetAll struct {
	Count int       `json:"count"`
	Users []UserDTO `json:"users"`
}
