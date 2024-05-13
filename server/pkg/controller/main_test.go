package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kirkalyn13/xyz-books-app/server/pkg/db"
)

var (
	gormFile   = "gorm.db"
	testFile   = "gorm-test.db"
	sourcePath = filepath.Join(".", "test-fixtures", testFile)
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	setup()
	exitCode := m.Run()
	defer teardown()

	os.Exit(exitCode)
}

func router() *gin.Engine {
	r := gin.Default()

	// GET Requests
	r.GET("/api/v1/books", GetBooks)
	r.GET("/api/v1/books/:id", GetBookByID)
	r.GET("/api/v1/books/isbn13/:isbn13", GetBookByISBN13)
	r.GET("/api/v1/authors", GetAuthors)
	r.GET("/api/v1/authors/:id", GetAuthorByID)
	r.GET("/api/v1/publishers", GetPublishers)
	r.GET("/api/v1/publishers/:id", GetPublisherByID)

	// POST Requests
	r.POST("/api/v1/books", AddBook)
	r.POST("/api/v1/authors", AddAuthor)
	r.POST("/api/v1/publishers", AddPublisher)

	// PUT Requests
	r.PUT("/api/v1/books/:id", EditBook)
	r.PUT("/api/v1/authors/:id", EditAuthor)
	r.PUT("/api/v1/publishers/:id", EditPublisher)

	// DELETE Requests
	r.DELETE("/api/v1/books/:id", DeleteBook)
	r.DELETE("/api/v1/authors/:id", DeleteAuthor)
	r.DELETE("/api/v1/publishers/:id", DeletePublisher)

	return r
}

func setup() {
	err := testDatabase()

	if err != nil {
		log.Fatalf("Error when setting up test database: %v", err)
	}

	err = db.LoadDatabase()

	if err != nil {
		log.Fatalf("Error when loading database: %v", err)
	}
}

func teardown() {
	err := os.Remove(gormFile)

	if err != nil {
		log.Fatalf("Error deleting test database: %v", err)
	}
}

func testDatabase() error {
	if _, err := os.Stat(gormFile); !os.IsNotExist(err) {
		teardown()
	}

	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(gormFile)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	err = destinationFile.Sync()
	if err != nil {
		return err
	}

	return nil
}

func makeRequest(method, url string, body interface{}) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(body)
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	writer := httptest.NewRecorder()
	router().ServeHTTP(writer, request)
	return writer
}
