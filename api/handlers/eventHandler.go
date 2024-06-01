package handlers

import (
	"net/http"

	"github.com/devlongs/ethereum-event-listener/api/models"
	"github.com/devlongs/ethereum-event-listener/database"
	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	DB database.Database
}

func (h *EventHandler) GetEvents(c *gin.Context) {
	var events []models.Event
	err := h.DB.GetEvents(&events)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}
