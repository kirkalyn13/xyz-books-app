package repository

import (
	"github.com/kirkalyn13/xyz-books-app/pkg/db"
	"github.com/kirkalyn13/xyz-books-app/pkg/model"
)

// GetBooks fetches the list of Books from the database
func GetBooks(searchQuery string) ([]model.Book, error) {
	var books []model.Book

	result := db.DB.Preload("Authors").Where("isbn13 LIKE ?", "%"+searchQuery+"%").Find(&books)

	if result.Error != nil {
		return []model.Book{}, result.Error
	}

	return books, nil
}

// GetBookByISBN13 fetches a Book based on a given ID from the database
func GetBookByISBN13(isbn13 string) (model.Book, error) {
	var book model.Book

	result := db.DB.Preload("Authors").Where("isbn13 = ?", isbn13).Find(&book)

	if result.Error != nil {
		return model.Book{}, result.Error
	}

	return book, nil
}

// GetBookByID fetches a Book based on a given ISBN13 from the database
func GetBookByID(id string) (model.Book, error) {
	var book model.Book

	result := db.DB.Preload("Authors").Where("id = ?", id).Find(&book)

	if result.Error != nil {
		return model.Book{}, result.Error
	}

	return book, nil
}

// AddBook adds a new Book entity from the database
func AddBook(book model.Book) (model.Book, error) {
	result := db.DB.Create(&book)

	if result.Error != nil {
		return model.Book{}, result.Error
	}

	return book, nil
}

// EditBook edits a Book entity from the database
func EditBook(book model.Book, id string) (model.Book, error) {
	result := db.DB.Where("id = ?", id).Save(&book)

	if result.Error != nil {
		return model.Book{}, result.Error
	}

	err := db.DB.Model(&book).Association("Authors").Replace(book.Authors)

	if err != nil {
		return model.Book{}, err
	}

	return book, nil
}

// DeleteBook deletes a Book entity from the database
func DeleteBook(id string) error {
	result := db.DB.Where("id = ?", id).Delete(&model.Book{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
