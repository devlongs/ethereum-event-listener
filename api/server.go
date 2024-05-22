package api

import (
	"github.com/devlongs/ethereum-event-listener/api/routes"

	"github.com/gin-gonic/gin"
)

func StartServer() {
    router := gin.Default()
    routes.SetupRoutes(router)
    router.Run(":8080")
}