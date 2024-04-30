package repository

import (
	"github.com/kirkalyn13/xyz-books-app/pkg/db"
	"github.com/kirkalyn13/xyz-books-app/pkg/model"
)

// GetBooks fetches the list of books from the database
func GetBooks() []model.Book {
	var books []model.Book

	db.DB.Unscoped().Find(&books)

	return books
}

// GetBooksByISBN13 fetches a book based on a given ISBN13 from the database
func GetBookByISBN13(isbn13 string) model.Book {
	var book model.Book

	db.DB.Unscoped().Find(&book, isbn13)

	return book
}

// AddBook adds a new book entity from the database
func AddBook(book model.Book) (model.Book, error) {
	result := db.DB.Unscoped().Create(&book)

	if result.Error != nil {
		return model.Book{}, result.Error
	}

	return book, nil
}

// EditBook edits a book entity from the database
func EditBook(book model.Book, isbn13 string) model.Book {
	db.DB.First(&book, isbn13)

	return book
}

// DeleteBook deletes a book entity from the database
func DeleteBook(isbn13 string) {
	db.DB.Delete(&model.Book{}, isbn13)
}
