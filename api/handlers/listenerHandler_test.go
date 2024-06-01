package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegisterListener(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockDB := &MockDatabase{}
	handler := &ListenerHandler{DB: mockDB}

	router.POST("/listeners", handler.RegisterListener)

	listenerJSON := `{"contract_address": "0x123", "event_name": "TestEvent"}`
	req, _ := http.NewRequest("POST", "/listeners", bytes.NewBuffer([]byte(listenerJSON)))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Listener registered")
}
