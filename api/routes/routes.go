package routes

import (
	"github.com/devlongs/ethereum-event-listener/api/handlers"
	"github.com/devlongs/ethereum-event-listener/api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
    api := router.Group("/api")
    {
        api.Use(middlewares.AuthMiddleware())
        api.POST("/listeners", handlers.RegisterListener)
        api.GET("/events", handlers.GetEvents)
    }
}