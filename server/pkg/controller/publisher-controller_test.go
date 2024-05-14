package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kirkalyn13/xyz-books-app/server/pkg/model"
	"github.com/stretchr/testify/assert"
)

var (
	testPublisher = model.Publisher{Name: "Test Publisher"}
)

func TestGetPublishersSuccess(t *testing.T) {
	router := router()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/publishers", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetPublishersFilterSuccess(t *testing.T) {
	router := router()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/publishers?q=publisher", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetPublisherByIDSuccess(t *testing.T) {
	router := router()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/publishers/1", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetPublisherByIDNotFound(t *testing.T) {
	router := router()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/publishers/12345", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestAddPublisherSuccess(t *testing.T) {
	router := router()

	reader, _ := structToReader(testPublisher)
	req, err := http.NewRequest(http.MethodPost, "/api/v1/publishers", reader)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestEditPublisherSuccess(t *testing.T) {
	router := router()

	reader, _ := structToReader(testPublisher)
	req, err := http.NewRequest(http.MethodPut, "/api/v1/publishers/4", reader)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestEditPublisherNotFound(t *testing.T) {
	router := router()

	reader, _ := structToReader(testPublisher)
	req, err := http.NewRequest(http.MethodPut, "/api/v1/publishers/12345", reader)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeletePublisherSuccess(t *testing.T) {
	router := router()

	req, err := http.NewRequest(http.MethodDelete, "/api/v1/publishers/4", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestDeletePublisherNotFound(t *testing.T) {
	router := router()

	req, err := http.NewRequest(http.MethodDelete, "/api/v1/publishers/12345", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
