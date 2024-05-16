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
	var res model.AuthorsResponse
	err = json.NewDecoder(w.Body).Decode(&res)
	assert.NoError(t, err)
	assert.Equal(t, 7, len(res.Authors))
}

func TestGetAuthorsFilteredSuccess(t *testing.T) {
	router := router()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/authors?q=king", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Must have 1 expected author
	var res model.AuthorsResponse
	err = json.NewDecoder(w.Body).Decode(&res)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(res.Authors))
}

func TestGetAuthorByIDSuccess(t *testing.T) {
	router := router()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/authors/5", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var res model.AuthorResponse
	err = json.NewDecoder(w.Body).Decode(&res)
	assert.NoError(t, err)
	assert.Equal(t, uint(5), res.Author.ID)
	assert.Equal(t, "Fannie", res.Author.FirstName)
	assert.Equal(t, "Peters", res.Author.MiddleName)
	assert.Equal(t, "Flagg", res.Author.LastName)
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
	req.Header.Set("Content-Type", "application/json")
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	var res model.AuthorResponse
	err = json.Unmarshal(w.Body.Bytes(), &res)
	assert.NoError(t, err)

	assert.Equal(t, testAuthor.FirstName, res.Author.FirstName)
	assert.Equal(t, testAuthor.LastName, res.Author.LastName)
	assert.Equal(t, testAuthor.MiddleName, res.Author.MiddleName)
}

func TestEditAuthorSuccess(t *testing.T) {
	router := router()

	reader, _ := structToReader(testAuthor)
	req, err := http.NewRequest(http.MethodPut, "/api/v1/authors/5", reader)
	req.Header.Set("Content-Type", "application/json")
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var res model.AuthorResponse
	err = json.Unmarshal(w.Body.Bytes(), &res)
	assert.NoError(t, err)

	assert.Equal(t, testAuthor.FirstName, res.Author.FirstName)
	assert.Equal(t, testAuthor.LastName, res.Author.LastName)
	assert.Equal(t, testAuthor.MiddleName, res.Author.MiddleName)
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
