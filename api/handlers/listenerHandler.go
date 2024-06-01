package handlers

import (
	"net/http"

	"github.com/devlongs/ethereum-event-listener/api/models"
	"github.com/devlongs/ethereum-event-listener/database"
	"github.com/devlongs/ethereum-event-listener/ethereum"
	"github.com/gin-gonic/gin"
)

type ListenerHandler struct {
	DB database.Database
}

func (h *ListenerHandler) RegisterListener(c *gin.Context) {
	var listener models.Listener
	if err := c.ShouldBindJSON(&listener); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.DB.AddListener(&listener)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	go ethereum.ListenForEvents(listener, h.DB)

	c.JSON(http.StatusOK, gin.H{"message": "Listener registered"})
}
