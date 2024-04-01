package database

import (
	"database/sql"
	"log"

	"github.com/YaroslavMiloslavsky/go-rest-api/api"
	"github.com/YaroslavMiloslavsky/go-rest-api/pkg/utils"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresDB struct{}

func CreateNewPostgresDB() *PostgresDB {
	return &PostgresDB{}
}

func (p *PostgresDB) GetAllUsers() (*[]api.UserFromDB, error) {
	connectionString := utils.GetPostgresUrl()

	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		log.Printf("Could not connect for postgres: %s\n", err)
		return nil, err
	}
	log.Println("Connection with db was established")
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Printf("Could not query for data: %s\n", err)
		return nil, err
	}
	defer rows.Close()
	log.Println("Query was executed and returned with results")
	var users []api.UserFromDB

	for rows.Next() {
		var user api.UserFromDB
		if err = rows.Scan(&user.Id, &user.Username, &user.Age, &user.Created, &user.Updated); err != nil {
			log.Println("Could not parse entry")
			continue
		}
		users = append(users, user)
	}

	log.Printf("Returning %d results\n", len(users))
	log.Printf("Closing db connection")
	return &users, nil
}
