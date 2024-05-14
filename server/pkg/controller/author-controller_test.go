package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kirkalyn13/xyz-books-app/server/pkg/model"
	"github.com/stretchr/testify/assert"
)

var (
	testAuthor = model.Author{
		FirstName:  "First Name",
		MiddleName: "Middle Name",
		LastName:   "Last Name",
	}
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
	router := router()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/authors/1", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	req, err = http.NewRequest(http.MethodGet, "/api/v1/authors/12345", nil)
	assert.NoError(t, err)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestAddAuthor(t *testing.T) {
	router := router()

	reader, _ := structToReader(testAuthor)
	req, err := http.NewRequest(http.MethodPost, "/api/v1/authors", reader)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestEditAuthor(t *testing.T) {
	router := router()

	reader, _ := structToReader(testAuthor)
	req, err := http.NewRequest(http.MethodPut, "/api/v1/authors/5", reader)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	req, err = http.NewRequest(http.MethodPut, "/api/v1/authors/12345", reader)
	assert.NoError(t, err)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeleteAuthor(t *testing.T) {
	router := router()

	req, err := http.NewRequest(http.MethodDelete, "/api/v1/authors/6", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)

	req, err = http.NewRequest(http.MethodDelete, "/api/v1/authors/12345", nil)
	assert.NoError(t, err)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
