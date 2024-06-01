package api

import (
	"github.com/devlongs/ethereum-event-listener/api/handlers"
	"github.com/devlongs/ethereum-event-listener/api/routes"
	"github.com/devlongs/ethereum-event-listener/config"
	"github.com/devlongs/ethereum-event-listener/database"
	"github.com/gin-gonic/gin"
)

type Server interface {
	Start() error
	Stop() error
}

type GinServer struct {
	Engine *gin.Engine
	DB     database.Database
}

func (s *GinServer) Start() error {
	eventHandler := &handlers.EventHandler{DB: s.DB}
	listenerHandler := &handlers.ListenerHandler{DB: s.DB}
	routes.SetupRoutes(s.Engine, eventHandler, listenerHandler)
	return s.Engine.Run(":8080")
}

func (s *GinServer) Stop() error {
	// Implement graceful shutdown if necessary
	return nil
}

func StartServer(config *config.Config, db database.Database) Server {
	router := gin.Default()
	server := &GinServer{Engine: router, DB: db}
	go server.Start()
	return server
}
