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
	testAuthor = model.Author{
		FirstName:  "First Name",
		MiddleName: "Middle Name",
		LastName:   "Last Name",
	}
)

func TestGetAuthorsSuccess(t *testing.T) {
	router := router()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/authors", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Must have 7 expected authors
	var results AuthorsResponse
	err = json.NewDecoder(w.Body).Decode(&results)
	assert.NoError(t, err)
	assert.Equal(t, 7, len(results.Authors))
}

func TestGetAuthorsFilteredSuccess(t *testing.T) {
	router := router()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/authors?q=king", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Must have 1 expected author
	var results AuthorsResponse
	err = json.NewDecoder(w.Body).Decode(&results)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(results.Authors))
}

func TestGetAuthorByIDSuccess(t *testing.T) {
	router := router()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/authors/5", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var result AuthorResponse
	err = json.NewDecoder(w.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, uint(5), result.Author.ID)
	assert.Equal(t, "Fannie", result.Author.FirstName)
	assert.Equal(t, "Peters", result.Author.MiddleName)
	assert.Equal(t, "Flagg", result.Author.LastName)
}

func TestGetAuthorByIDNotFound(t *testing.T) {
	router := router()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/authors/12345", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestAddAuthorSuccess(t *testing.T) {
	router := router()

	reader, _ := structToReader(testAuthor)
	req, err := http.NewRequest(http.MethodPost, "/api/v1/authors", reader)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestEditAuthorSuccess(t *testing.T) {
	router := router()

	reader, _ := structToReader(testAuthor)
	req, err := http.NewRequest(http.MethodPut, "/api/v1/authors/5", reader)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestEditAuthorNotFound(t *testing.T) {
	router := router()

	reader, _ := structToReader(testAuthor)
	req, err := http.NewRequest(http.MethodPut, "/api/v1/authors/12345", reader)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeleteAuthorSuccess(t *testing.T) {
	router := router()

	req, err := http.NewRequest(http.MethodDelete, "/api/v1/authors/6", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)

	req, err = http.NewRequest(http.MethodGet, "/api/v1/authors/6", nil)
	assert.NoError(t, err)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeleteAuthorNotFound(t *testing.T) {
	router := router()

	req, err := http.NewRequest(http.MethodDelete, "/api/v1/authors/12345", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
