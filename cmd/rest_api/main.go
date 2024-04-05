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
	log      = logger.CreateNew()
	testData = flag.Bool("testing", false, "Populate data for testing")
	dropDB   = flag.Bool("drop-tables", false, "Drop all DB tables")
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

func onStartUp() {
	flag.Parse()
	connectionString := utils.GetPostgresUrl()

	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		log.Printf("Could not connect for postgres: %s\n", err)
		return
	}
	defer db.Close()

	scriptsToExecute := make([][]byte, 8)
	scriptsToExecute = append(scriptsToExecute, dropDb())
	scriptsToExecute = append(scriptsToExecute, startupDb())
	scriptsToExecute = append(scriptsToExecute, populateTestData())

	for _, script := range scriptsToExecute {
		if script != nil {
			executeQuery(db, script)
		}
	}
}

func populateTestData() []byte {
	if *testData {
		sqlBytes, err := schemaSQL.ReadFile("sql/test_data.sql")
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Testing data was populated")
		return sqlBytes
	}
	return nil
}

func startupDb() []byte {
	sqlBytes, err := schemaSQL.ReadFile("sql/startup.sql")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	log.Println("Tables created successfully!")
	return sqlBytes
}

func dropDb() []byte {
	if *dropDB {
		sqlBytes, err := schemaSQL.ReadFile("sql/drop_db.sql")
		if err != nil {
			log.Fatal(err)
		}
		log.Println("DB tables were dropped")
		return sqlBytes
	}
	return nil
} 

func executeQuery(db *sql.DB, sqlBytes[]byte) {
	queries := string(sqlBytes)
	_, err := db.Exec(queries)
	if err != nil {
		log.Fatal(err)
	}
}
