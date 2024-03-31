package utils

import (
	"fmt"
	"os"

	"github.com/YaroslavMiloslavsky/go-rest-api/pkg/db_const"
	"github.com/YaroslavMiloslavsky/go-rest-api/pkg/server_const"
)

func GetServerURL() string {
	host := os.Getenv(server_const.EnvHostKey)
	port := os.Getenv(server_const.EnvPortKey)
	return host + ":" + port
}

func GetPostgresUrl(dbName string) string {
	user := os.Getenv(db_const.EnvPostgresUsernameKey)
	password := os.Getenv(db_const.EnvPostgresPasswordKey)
	host := os.Getenv(db_const.EnvPostgresHostKey)
	port := os.Getenv(db_const.EnvPostgresPortKey)

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbName)
}
