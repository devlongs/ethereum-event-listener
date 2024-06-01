package database

import (
	"database/sql"
	"fmt"

	"github.com/devlongs/ethereum-event-listener/api/models"
	_ "github.com/go-sql-driver/mysql"
)

type SQLDatabase struct {
	DB     *sql.DB
	config map[string]string
}

func (db *SQLDatabase) Configure(config map[string]string) {
	db.config = config
}

func (db *SQLDatabase) Connect() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		db.config["DB_USER"], db.config["DB_PASSWORD"], db.config["DB_HOST"], db.config["DB_PORT"], db.config["DB_NAME"])

	database, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	db.DB = database
	return nil
}

func (db *SQLDatabase) Disconnect() error {
	return db.DB.Close()
}

func (db *SQLDatabase) GetEvents(events *[]models.Event) error {
	rows, err := db.DB.Query("SELECT * FROM events")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var event models.Event
		if err := rows.Scan(&event.ContractAddress, &event.EventName, &event.BlockNumber, &event.TransactionHash, &event.Data); err != nil {
			return err
		}
		*events = append(*events, event)
	}
	return nil
}

func (db *SQLDatabase) AddListener(listener *models.Listener) error {
	query := `INSERT INTO listeners (contract_address, event_name) VALUES (?, ?)`
	_, err := db.DB.Exec(query, listener.ContractAddress, listener.EventName)
	return err
}

func (db *SQLDatabase) AddEvent(event *models.Event) error {
	query := `INSERT INTO events (contract_address, event_name, block_number, transaction_hash, data) VALUES (?, ?, ?, ?, ?)`
	_, err := db.DB.Exec(query, event.ContractAddress, event.EventName, event.BlockNumber, event.TransactionHash, event.Data)
	return err
}
