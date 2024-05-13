package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAuthors(t *testing.T) {
	router := router()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/authors", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetAuthorByID(t *testing.T) {
	_, err := http.NewRequest(http.MethodGet, "/api/v1/authors/1", nil)
	assert.NoError(t, err)
}

func TestAddAuthor(t *testing.T) {
	_, err := http.NewRequest(http.MethodPost, "/api/v1/authors", nil)
	assert.NoError(t, err)
}

func TestEditAuthor(t *testing.T) {
	_, err := http.NewRequest(http.MethodPut, "/api/v1/authors/1", nil)
	assert.NoError(t, err)
}

func TestDeleteAuthor(t *testing.T) {
	_, err := http.NewRequest(http.MethodDelete, "/api/v1/authors/1", nil)
	assert.NoError(t, err)
}
