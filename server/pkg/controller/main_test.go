package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kirkalyn13/xyz-books-app/server/pkg/db"
	"github.com/kirkalyn13/xyz-books-app/server/pkg/model"
)

var (
	gormFile   = "gorm.db"
	testFile   = "gorm-test.db"
	sourcePath = filepath.Join(".", "test-fixtures", testFile)
)

// Response types
type (
	BooksResponse struct {
		Books []model.Book `json:"books,omitempty"`
	}
	AuthorsResponse struct {
		Authors []model.Author `json:"authors,omitempty"`
	}
	PublishersResponse struct {
		Publishers []model.Publisher `json:"publishers,omitempty"`
	}
	BookResponse struct {
		Book model.Book `json:"book,omitempty"`
	}
	AuthorResponse struct {
		Author model.Author `json:"author,omitempty"`
	}
	PublisherResponse struct {
		Publisher model.Publisher `json:"publisher,omitempty"`
	}
)

// TestMain runs API functional tests main code
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	setup()
	exitCode := m.Run()
	defer teardown()

	os.Exit(exitCode)
}

// router setups router for functional tests
// Include controller endpoints for API to be tested here
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

// setup setups the database for testing
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

// teardown removes test database
func teardown() {
	err := os.Remove(gormFile)

	if err != nil {
		log.Fatalf("Error deleting test database: %v", err)
	}
}

// testDatabase initializes test database
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

// structToReader converts test data to input io reader for test requests
func structToReader(data interface{}) (io.Reader, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(jsonData)
	return reader, nil
}
