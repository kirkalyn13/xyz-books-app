package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kirkalyn13/xyz-books-app/server/pkg/model"
	"github.com/stretchr/testify/assert"
)

var (
	invalidBook = model.Book{ISBN13: "123456789"}
	testBook    = model.Book{
		Title:           "Test Book",
		ISBN13:          "9780547928227",
		ISBN10:          "054792822X",
		ListPrice:       1137,
		PublicationYear: 2020,
		ImageURL:        "",
		Edition:         "First Edition",
	}
)

func TestGetBooks(t *testing.T) {
	router := router()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/books", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetBookByISBN13Success(t *testing.T) {
	router := router()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/books/isbn13/9781891830853", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetBookByISBN13InvalidISBN13(t *testing.T) {
	router := router()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/books/isbn13/123456789", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetBookByISBN13NotFound(t *testing.T) {
	router := router()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/books/isbn13/9780547928227", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetBookByIDSuccess(t *testing.T) {
	router := router()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/books/1", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAddBookIDSuccess(t *testing.T) {
	// router := router()

	// reader, _ := structToReader(testBook)
	// req, err := http.NewRequest(http.MethodPost, "/api/v1/books", reader)
	// assert.NoError(t, err)

	// w := httptest.NewRecorder()
	// router.ServeHTTP(w, req)
	// assert.Equal(t, http.StatusCreated, w.Code)
}

func TestAddBookInvalidISBN(t *testing.T) {
	router := router()

	reader, _ := structToReader(invalidBook)
	req, err := http.NewRequest(http.MethodPost, "/api/v1/books", reader)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAddBookNotNoISBN(t *testing.T) {
	router := router()

	testBook.ISBN13 = ""
	testBook.ISBN10 = ""

	reader, _ := structToReader(testBook)
	req, err := http.NewRequest(http.MethodPost, "/api/v1/books", reader)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestEditBookSuccess(t *testing.T) {
	// router := router()

	// reader, _ := structToReader(testBook)
	// req, err := http.NewRequest(http.MethodPut, "/api/v1/books/5", reader)
	// assert.NoError(t, err)

	// w := httptest.NewRecorder()
	// router.ServeHTTP(w, req)
	// assert.Equal(t, http.StatusOK, w.Code)
}

func TestEditBookNotFound(t *testing.T) {
	// router := router()

	// reader, _ := structToReader(testBook)
	// req, err := http.NewRequest(http.MethodPut, "/api/v1/books/12345", reader)
	// assert.NoError(t, err)

	// w := httptest.NewRecorder()
	// router.ServeHTTP(w, req)
	// assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeleteBookSuccess(t *testing.T) {
	router := router()

	req, err := http.NewRequest(http.MethodDelete, "/api/v1/books/5", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestDeleteBookNotFound(t *testing.T) {
	router := router()

	req, err := http.NewRequest(http.MethodDelete, "/api/v1/books/12345", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
