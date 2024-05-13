package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBooks(t *testing.T) {
	router := router()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/books", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetBookByISBN13(t *testing.T) {
	router := router()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/books/isbn13/9781891830853", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	req, err = http.NewRequest(http.MethodGet, "/api/v1/books/isbn13/123456789", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetBookByID(t *testing.T) {
	_, err := http.NewRequest(http.MethodGet, "/api/v1/books/1", nil)
	assert.NoError(t, err)
}

func TestAddBook(t *testing.T) {
	_, err := http.NewRequest(http.MethodPost, "/api/v1/books", nil)
	assert.NoError(t, err)
}

func TestEditBook(t *testing.T) {
	_, err := http.NewRequest(http.MethodPut, "/api/v1/books/1", nil)
	assert.NoError(t, err)
}

func TestDeleteBook(t *testing.T) {
	_, err := http.NewRequest(http.MethodDelete, "/api/v1/books/1", nil)
	assert.NoError(t, err)
}
