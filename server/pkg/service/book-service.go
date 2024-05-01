package service

import (
	"github.com/kirkalyn13/xyz-books-app/pkg/model"
	"github.com/kirkalyn13/xyz-books-app/pkg/repository"
)

// GetBooks fetches the list of Books
func GetBooks(searchQuery string) ([]model.Book, error) {
	books, err := repository.GetBooks(searchQuery)

	if err != nil {
		return []model.Book{}, err
	}

	return books, nil
}

// GetBooksByISBN13 fetches a Book based on a given ISBN13
func GetBookByISBN13(isbn13 string) (model.Book, error) {
	book, err := repository.GetBookByISBN13(isbn13)

	if err != nil {
		return model.Book{}, err
	}

	return book, nil
}

// AddBook adds a new Book entity
func AddBook(book model.Book) (model.Book, error) {
	result, err := repository.AddBook(book)

	if err != nil {
		return model.Book{}, err
	}

	return result, nil
}

// EditBook edits a Book entity
func EditBook(book model.Book, isbn13 string) (model.Book, error) {
	book, err := repository.EditBook(book, isbn13)

	if err != nil {
		return model.Book{}, err
	}

	return book, nil
}

// DeleteBook deletes a Book entity
func DeleteBook(isbn13 string) error {
	err := repository.DeleteBook(isbn13)

	if err != nil {
		return err
	}

	return nil
}
