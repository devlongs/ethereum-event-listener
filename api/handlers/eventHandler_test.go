package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/devlongs/ethereum-event-listener/api/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockDatabase struct{}

func (m *MockDatabase) Configure(config map[string]string) {}
func (m *MockDatabase) Connect() error                     { return nil }
func (m *MockDatabase) Disconnect() error                  { return nil }
func (m *MockDatabase) GetEvents(events *[]models.Event) error {
	*events = append(*events, models.Event{ContractAddress: "0x123", EventName: "TestEvent"})
	return nil
}
func (m *MockDatabase) AddListener(listener *models.Listener) error { return nil }
func (m *MockDatabase) AddEvent(event *models.Event) error          { return nil }

func TestGetEvents(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockDB := &MockDatabase{}
	handler := &EventHandler{DB: mockDB}

	router.GET("/events", handler.GetEvents)

	req, _ := http.NewRequest("GET", "/events", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "TestEvent")
}
