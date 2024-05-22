package handlers

import (
	"net/http"

	"github.com/devlongs/ethereum-event-listener/api/models"
	"github.com/devlongs/ethereum-event-listener/database"

	"github.com/gin-gonic/gin"
)

func GetEvents(c *gin.Context) {
    var events []models.Event
    if err := database.DB.Find(&events).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, events)
}