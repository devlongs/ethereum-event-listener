package database

import (
	"github.com/devlongs/ethereum-event-listener/api/models"
)

type Database interface {
	Configure(config map[string]string)
	Connect() error
	Disconnect() error
	GetEvents(events *[]models.Event) error
	AddListener(listener *models.Listener) error
	AddEvent(event *models.Event) error
}
