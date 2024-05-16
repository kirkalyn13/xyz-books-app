package controller

import (
	"encoding/json"
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

	// Must have 4 expected publishers
	var res model.PublishersResponse
	err = json.NewDecoder(w.Body).Decode(&res)
	assert.NoError(t, err)
	assert.Equal(t, 4, len(res.Publishers))
}

func TestGetPublishersFilterSuccess(t *testing.T) {
	router := router()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/publishers?q=publisher", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Must have 1 expected publisher
	var res model.PublishersResponse
	err = json.NewDecoder(w.Body).Decode(&res)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(res.Publishers))
}

func TestGetPublisherByIDSuccess(t *testing.T) {
	router := router()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/publishers/1", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var res model.PublisherResponse
	err = json.NewDecoder(w.Body).Decode(&res)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), res.Publisher.ID)
	assert.Equal(t, "Paste Magazine", res.Publisher.Name)
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
	req.Header.Set("Content-Type", "application/json")
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	var res model.PublisherResponse
	err = json.Unmarshal(w.Body.Bytes(), &res)
	assert.NoError(t, err)

	assert.Equal(t, testPublisher.Name, res.Publisher.Name)
}

func TestEditPublisherSuccess(t *testing.T) {
	router := router()
	testPublisher.Name = "Edited Publisher"

	reader, _ := structToReader(testPublisher)
	req, err := http.NewRequest(http.MethodPut, "/api/v1/publishers/4", reader)
	req.Header.Set("Content-Type", "application/json")
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var res model.PublisherResponse
	err = json.Unmarshal(w.Body.Bytes(), &res)
	assert.NoError(t, err)

	assert.Equal(t, testPublisher.Name, res.Publisher.Name)
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

	req, err = http.NewRequest(http.MethodGet, "/api/v1/publishers/4", nil)
	assert.NoError(t, err)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeletePublisherNotFound(t *testing.T) {
	router := router()

	req, err := http.NewRequest(http.MethodDelete, "/api/v1/publishers/12345", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
