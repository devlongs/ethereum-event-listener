package main

import (
	"log"

	"github.com/devlongs/ethereum-event-listener/api"
	"github.com/devlongs/ethereum-event-listener/config"
	"github.com/devlongs/ethereum-event-listener/database"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	var db database.Database
	dbConfig := map[string]string{
		"DB_HOST":     cfg.DBHost,
		"DB_PORT":     cfg.DBPort,
		"DB_USER":     cfg.DBUser,
		"DB_PASSWORD": cfg.DBPassword,
		"DB_NAME":     cfg.DBName,
		"DB_URI":      cfg.DBHost,
	}

	switch cfg.DBType {
	case "sql":
		db = &database.SQLDatabase{}
	case "nosql":
		db = &database.NoSQLDatabase{}
	default:
		log.Fatal("Unsupported database type")
	}

	db.Configure(dbConfig)

	err = db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Disconnect()

	server := api.StartServer(cfg, db)
	defer server.Stop()
}
