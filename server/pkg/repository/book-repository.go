package repository

import (
	"github.com/kirkalyn13/xyz-books-app/pkg/db"
	"github.com/kirkalyn13/xyz-books-app/pkg/model"
)

// GetBooks fetches the list of Books from the database
func GetBooks(searchQuery string) ([]model.Book, error) {
	var books []model.Book

	result := db.DB.Where("isbn13 LIKE ?", "%"+searchQuery+"%").Find(&books)

	if result.Error != nil {
		return []model.Book{}, result.Error
	}

	return books, nil
}

// GetBooksByISBN13 fetches a Book based on a given ISBN13 from the database
func GetBookByISBN13(isbn13 string) (model.Book, error) {
	var book model.Book

	result := db.DB.Where("isbn13 = ?", isbn13).Find(&book)

	if result.Error != nil {
		return model.Book{}, result.Error
	}

	return book, nil
}

// AddBook adds a new Book entity from the database
func AddBook(book model.Book) (model.Book, error) {
	result := db.DB.Unscoped().Create(&book)

	if result.Error != nil {
		return model.Book{}, result.Error
	}

	return book, nil
}

// EditBook edits a Book entity from the database
func EditBook(book model.Book, isbn13 string) (model.Book, error) {
	result := db.DB.Where("isbn13 = ?", isbn13).Updates(&book)

	if result.Error != nil {
		return model.Book{}, result.Error
	}

	return book, nil
}

// DeleteBook deletes a Book entity from the database
func DeleteBook(isbn13 string) error {
	result := db.DB.Where("isbn13 = ?", isbn13).Delete(&model.Book{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
