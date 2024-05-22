package main

import (
	"log"

	"github.com/devlongs/ethereum-event-listener/api"
	"github.com/devlongs/ethereum-event-listener/database"
)

func main() {
    if err := database.ConnectDB(); err != nil {
        log.Fatal(err)
    }

    api.StartServer()
}