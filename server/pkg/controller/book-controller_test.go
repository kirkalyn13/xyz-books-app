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
	publisherID = uint(1)
	testBook    = model.Book{
		ID:              6,
		Title:           "Test Book",
		ISBN13:          "9780547928227",
		ISBN10:          "054792822X",
		ListPrice:       1137,
		PublicationYear: 2020,
		ImageURL:        "",
		Edition:         "First Edition",
		PublisherID:     &publisherID,
		Authors: []*model.Author{
			{ID: 1},
		},
	}
)

func TestGetBooksSuccess(t *testing.T) {
	r := router()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/books", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Must have 5 expected books
	var res model.BooksResponse
	err = json.NewDecoder(w.Body).Decode(&res)
	assert.NoError(t, err)
	assert.Equal(t, 5, len(res.Books))
}

func TestGetBooksFilterSuccess(t *testing.T) {
	r := router()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/books?q=97816", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Must have 3 expected books
	var res model.BooksResponse
	err = json.NewDecoder(w.Body).Decode(&res)
	assert.NoError(t, err)
	assert.Equal(t, 3, len(res.Books))
}

func TestGetBookByISBN13Success(t *testing.T) {
	r := router()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/books/isbn13/9781891830853", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var res model.BookResponse
	err = json.NewDecoder(w.Body).Decode(&res)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), res.Book.ID)
	assert.Equal(t, "American Elf", res.Book.Title)
	assert.Equal(t, "9781891830853", res.Book.ISBN13)
	assert.Equal(t, "1891830856", res.Book.ISBN10)
	assert.Equal(t, int(1000), int(res.Book.ListPrice))
	assert.Equal(t, int(2004), int(res.Book.PublicationYear))
	assert.Equal(t, "https://images-na.ssl-images-amazon.com/images/S/compressed.photo.goodreads.com/books/1343244815i/15770036.jpg", res.Book.ImageURL)
	assert.Equal(t, "Book 2", res.Book.Edition)
	assert.Equal(t, uint(1), *res.Book.PublisherID)
	assert.Equal(t, 3, len(res.Book.Authors))
	assert.Equal(t, uint(1), res.Book.Publisher.ID)
	assert.Equal(t, "Paste Magazine", res.Book.Publisher.Name)
}

func TestGetBookByISBN13InvalidISBN13(t *testing.T) {
	r := router()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/books/isbn13/123456789", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetBookByISBN13NotFound(t *testing.T) {
	r := router()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/books/isbn13/9780547928227", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetBookByIDSuccess(t *testing.T) {
	r := router()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/books/1", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var res model.BookResponse
	err = json.NewDecoder(w.Body).Decode(&res)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), res.Book.ID)
	assert.Equal(t, "American Elf", res.Book.Title)
	assert.Equal(t, "9781891830853", res.Book.ISBN13)
	assert.Equal(t, "1891830856", res.Book.ISBN10)
	assert.Equal(t, int(1000), int(res.Book.ListPrice))
	assert.Equal(t, int(2004), int(res.Book.PublicationYear))
	assert.Equal(t, "https://images-na.ssl-images-amazon.com/images/S/compressed.photo.goodreads.com/books/1343244815i/15770036.jpg", res.Book.ImageURL)
	assert.Equal(t, "Book 2", res.Book.Edition)
	assert.Equal(t, uint(1), *res.Book.PublisherID)
	assert.Equal(t, 3, len(res.Book.Authors))
	assert.Equal(t, uint(1), res.Book.Publisher.ID)
	assert.Equal(t, "Paste Magazine", res.Book.Publisher.Name)
}

func TestAddBookIDSuccess(t *testing.T) {
	r := router()

	reader, _ := structToReader(testBook)
	req, err := http.NewRequest(http.MethodPost, "/api/v1/books", reader)
	req.Header.Set("Content-Type", "application/json")
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	var res model.BookResponse
	err = json.Unmarshal(w.Body.Bytes(), &res)
	assert.NoError(t, err)

	assert.Equal(t, testBook.Title, res.Book.Title)
	assert.Equal(t, testBook.ISBN13, res.Book.ISBN13)
	assert.Equal(t, testBook.ISBN10, res.Book.ISBN10)
	assert.Equal(t, testBook.ListPrice, res.Book.ListPrice)
	assert.Equal(t, testBook.PublicationYear, res.Book.PublicationYear)
	assert.Equal(t, testBook.Edition, res.Book.Edition)
}

func TestAddBookInvalidISBN13(t *testing.T) {
	r := router()
	testBook.ISBN13 = "12345XYZ"

	reader, _ := structToReader(testBook)
	req, err := http.NewRequest(http.MethodPost, "/api/v1/books", reader)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAddBookInvalidISBN10(t *testing.T) {
	r := router()
	testBook.ISBN10 = "12345XYZ"

	reader, _ := structToReader(testBook)
	req, err := http.NewRequest(http.MethodPost, "/api/v1/books", reader)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAddBookNoISBN(t *testing.T) {
	r := router()

	testBook.ISBN13 = ""
	testBook.ISBN10 = ""

	reader, _ := structToReader(testBook)
	req, err := http.NewRequest(http.MethodPost, "/api/v1/books", reader)
	req.Header.Set("Content-Type", "application/json")
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestEditBookSuccess(t *testing.T) {
	r := router()

	testBook.ID = 5
	testBook.ISBN13 = "9781593275846"
	testBook.ISBN10 = "1593275846"

	reader, _ := structToReader(testBook)
	req, err := http.NewRequest(http.MethodPut, "/api/v1/books/5", reader)
	req.Header.Set("Content-Type", "application/json")
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var res model.BookResponse
	err = json.Unmarshal(w.Body.Bytes(), &res)
	assert.NoError(t, err)

	assert.Equal(t, uint(5), res.Book.ID)
	assert.Equal(t, testBook.Title, res.Book.Title)
	assert.Equal(t, testBook.ISBN13, res.Book.ISBN13)
	assert.Equal(t, testBook.ISBN10, res.Book.ISBN10)
	assert.Equal(t, testBook.ListPrice, res.Book.ListPrice)
	assert.Equal(t, testBook.PublicationYear, res.Book.PublicationYear)
	assert.Equal(t, testBook.Edition, res.Book.Edition)
}

func TestEditBookNotFound(t *testing.T) {
	r := router()

	reader, _ := structToReader(testBook)
	req, err := http.NewRequest(http.MethodPut, "/api/v1/books/12345", reader)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeleteBookSuccess(t *testing.T) {
	r := router()

	req, err := http.NewRequest(http.MethodDelete, "/api/v1/books/5", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)

	req, err = http.NewRequest(http.MethodGet, "/api/v1/books/5", nil)
	assert.NoError(t, err)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeleteBookNotFound(t *testing.T) {
	r := router()

	req, err := http.NewRequest(http.MethodDelete, "/api/v1/books/12345", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetIncompleteISBNsSuccess(t *testing.T) {
	r := router()
	testBook.ISBN13 = ""

	reader, _ := structToReader(testBook)
	req, err := http.NewRequest(http.MethodPost, "/api/v1/books", reader)
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	req, err = http.NewRequest(http.MethodGet, "/api/v1/books/isbn/incomplete", nil)
	assert.NoError(t, err)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var res model.BooksResponse
	err = json.Unmarshal(w.Body.Bytes(), &res)
	assert.NoError(t, err)

	assert.Equal(t, 1, len(res.Books))
	assert.Equal(t, testBook.Title, res.Books[0].Title)
	assert.Equal(t, testBook.ISBN13, res.Books[0].ISBN13)
	assert.Equal(t, testBook.ISBN10, res.Books[0].ISBN10)
	assert.Equal(t, testBook.ListPrice, res.Books[0].ListPrice)
	assert.Equal(t, testBook.PublicationYear, res.Books[0].PublicationYear)
	assert.Equal(t, testBook.Edition, res.Books[0].Edition)
}
