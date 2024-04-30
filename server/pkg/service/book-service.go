package service

import (
	"github.com/kirkalyn13/xyz-books-app/pkg/model"
	"github.com/kirkalyn13/xyz-books-app/pkg/repository"
)

// GetBooks fetches the list of Books
func GetBooks() []model.Book {
	return repository.GetBooks()
}

// GetBooksByISBN13 fetches a Book based on a given ISBN13
func GetBookByISBN13(isbn13 string) model.Book {
	return repository.GetBookByISBN13(isbn13)
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
func EditBook(book model.Book, isbn13 string) model.Book {
	result := repository.EditBook(book, isbn13)

	return result
}

// DeleteBook deletes a Book entity
func DeleteBook(isbn13 string) {
	repository.DeleteBook(isbn13)
}
