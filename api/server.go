package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	Server *http.Server
}

func (s *GinServer) Start() error {
	eventHandler := &handlers.EventHandler{DB: s.DB}
	listenerHandler := &handlers.ListenerHandler{DB: s.DB}
	routes.SetupRoutes(s.Engine, eventHandler, listenerHandler)

	s.Server = &http.Server{
		Addr:    ":8080",
		Handler: s.Engine,
	}

	go func() {
		if err := s.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	return s.Stop()
}

func (s *GinServer) Stop() error {
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
	return nil
}

func StartServer(config *config.Config, db database.Database) Server {
	router := gin.Default()
	server := &GinServer{Engine: router, DB: db}
	go server.Start()
	return server
}
