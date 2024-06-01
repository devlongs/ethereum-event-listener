package routes

import (
	"github.com/devlongs/ethereum-event-listener/api/handlers"
	"github.com/devlongs/ethereum-event-listener/api/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, eventHandler *handlers.EventHandler, listenerHandler *handlers.ListenerHandler) {
	api := router.Group("/api")
	{
		api.Use(middlewares.AuthMiddleware())
		api.GET("/events", eventHandler.GetEvents)
		api.POST("/listeners", listenerHandler.RegisterListener)
	}
}
