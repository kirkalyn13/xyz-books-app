package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kirkalyn13/xyz-books-app/server/pkg/model"
	"github.com/stretchr/testify/assert"
)

var (
	testPublisher    = model.Publisher{Name: "Test Publisher"}
	invalidPublisher = model.Publisher{ID: 12345, Name: "Test Publisher"}
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
	router := router()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/publishers/1", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	req, err = http.NewRequest(http.MethodGet, "/api/v1/publishers/12345", nil)
	assert.NoError(t, err)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestAddPublisher(t *testing.T) {
	router := router()

	req, err := http.NewRequest(http.MethodPost, "/api/v1/publishers", nil)
	assert.NoError(t, err)

	reader, _ := structToReader(testPublisher)
	req, err = http.NewRequest(http.MethodPost, "/api/v1/publishers", reader)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestEditPublisher(t *testing.T) {
	router := router()

	reader, _ := structToReader(testPublisher)
	req, err := http.NewRequest(http.MethodPut, "/api/v1/publishers/4", reader)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// reader, _ = structToReader(invalidPublisher)
	// req, err = http.NewRequest(http.MethodPut, "/api/v1/publishers/12345", reader)
	// assert.NoError(t, err)

	// w = httptest.NewRecorder()
	// router.ServeHTTP(w, req)
	// assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeletePublisher(t *testing.T) {
	router := router()

	req, err := http.NewRequest(http.MethodDelete, "/api/v1/publishers/1", nil)
	assert.NoError(t, err)

	// w := httptest.NewRecorder()
	// router.ServeHTTP(w, req)
	// assert.Equal(t, http.StatusOK, w.Code)

	req, err = http.NewRequest(http.MethodDelete, "/api/v1/publishers/12345", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
