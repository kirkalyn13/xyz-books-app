package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPublishers(t *testing.T) {
	router := router()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/publishers", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetPublisherByID(t *testing.T) {
	_, err := http.NewRequest(http.MethodGet, "/api/v1/publishers/1", nil)
	assert.NoError(t, err)
}

func TestAddPublisher(t *testing.T) {
	_, err := http.NewRequest(http.MethodPost, "/api/v1/publishers", nil)
	assert.NoError(t, err)
}

func TestEditPublisher(t *testing.T) {
	_, err := http.NewRequest(http.MethodPut, "/api/v1/publishers/1", nil)
	assert.NoError(t, err)
}

func TestDeletePublisher(t *testing.T) {
	_, err := http.NewRequest(http.MethodDelete, "/api/v1/publishers/1", nil)
	assert.NoError(t, err)
}
