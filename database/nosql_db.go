package database

import (
	"context"
	"time"

	"github.com/devlongs/ethereum-event-listener/api/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NoSQLDatabase struct {
	Client *mongo.Client
	DB     *mongo.Database
	config map[string]string
}

func (db *NoSQLDatabase) Configure(config map[string]string) {
	db.config = config
}

func (db *NoSQLDatabase) Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(db.config["DB_URI"]))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	db.Client = client
	db.DB = client.Database(db.config["DB_NAME"])
	return nil
}

func (db *NoSQLDatabase) Disconnect() error {
	return db.Client.Disconnect(context.Background())
}

func (db *NoSQLDatabase) GetEvents(events *[]models.Event) error {
	collection := db.DB.Collection("events")
	cur, err := collection.Find(context.Background(), nil)
	if err != nil {
		return err
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var event models.Event
		err := cur.Decode(&event)
		if err != nil {
			return err
		}
		*events = append(*events, event)
	}
	return nil
}

func (db *NoSQLDatabase) AddListener(listener *models.Listener) error {
	collection := db.DB.Collection("listeners")
	_, err := collection.InsertOne(context.Background(), listener)
	return err
}

func (db *NoSQLDatabase) AddEvent(event *models.Event) error {
	collection := db.DB.Collection("events")
	_, err := collection.InsertOne(context.Background(), event)
	return err
}
