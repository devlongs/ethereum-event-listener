package handlers

import (
	"net/http"

	"github.com/devlongs/ethereum-event-listener/api/models"
	"github.com/devlongs/ethereum-event-listener/database"
	"github.com/devlongs/ethereum-event-listener/ethereum"

	"github.com/gin-gonic/gin"
)

func RegisterListener(c *gin.Context) {
    var listener models.Listener
    if err := c.ShouldBindJSON(&listener); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := database.DB.Create(&listener).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    go ethereum.ListenForEvents(listener)

    c.JSON(http.StatusOK, gin.H{"message": "Listener registered"})
}