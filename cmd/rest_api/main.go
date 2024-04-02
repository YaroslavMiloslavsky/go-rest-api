package main

import (
	"database/sql"
	"embed"
	"flag"
	"net/http"

	"github.com/YaroslavMiloslavsky/go-rest-api/internal/handler"
	"github.com/YaroslavMiloslavsky/go-rest-api/internal/logger"
	"github.com/YaroslavMiloslavsky/go-rest-api/internal/middleware"
	"github.com/YaroslavMiloslavsky/go-rest-api/pkg/utils"
)

//go:embed sql/*
var schemaSQL embed.FS

var (
	log = logger.CreateNew()
)

func main() {
	onStartUp()
	var router http.Handler = http.NewServeMux()
	handler.Handler(&router)
	mStack := middleware.CreateMiddlewareStack(
		middleware.StripSlashes,
	)

	serverURL := utils.GetServerURL()
	log.Println("about to serve " + serverURL)

	http.ListenAndServe(serverURL, mStack(router))
}

// For enabling testing properties for the app run executable with -testing flag
func onStartUp() {
	dummyData := flag.Bool("testing", false, "Populate data for testing")
	flag.Parse()
	connectionString := utils.GetPostgresUrl()

	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		log.Printf("Could not connect for postgres: %s\n", err)
		return
	}
	defer db.Close()

	sqlBytes, err := schemaSQL.ReadFile("sql/startup.sql")
	if err != nil {
		log.Fatal(err)
	}

	if *dummyData {
		sqlBytes, err = schemaSQL.ReadFile("sql/dummy_data.sql")
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Testing data was populated")
	}

	queries := string(sqlBytes)
	_, err = db.Exec(queries)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Tables created successfully!")
}
