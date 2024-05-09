package service

import (
	"errors"
	"strconv"

	"github.com/kirkalyn13/xyz-books-app/pkg/model"
	"github.com/kirkalyn13/xyz-books-app/pkg/mq"
	"github.com/kirkalyn13/xyz-books-app/pkg/repository"
)

// GetBooks fetches the list of Books
func GetBooks(searchQuery string) ([]model.Book, error) {
	bookResults := []model.Book{}
	books, err := repository.GetBooks(searchQuery)

	if err != nil {
		return []model.Book{}, err
	}

	for _, book := range books {
		publisher, _ := repository.GetPublisherByID(strconv.Itoa(int(*book.PublisherID)))
		publisher.Books = []*model.Book{}
		book.Publisher = publisher
		bookResults = append(bookResults, book)
	}

	return bookResults, nil
}

// GetBooksByISBN13 fetches a Book based on a given ISBN13
func GetBookByISBN13(isbn13 string) (model.Book, error) {
	book, err := repository.GetBookByISBN13(isbn13)

	if err != nil {
		return model.Book{}, err
	}

	publisher, _ := repository.GetPublisherByID(strconv.Itoa(int(*book.PublisherID)))
	publisher.Books = []*model.Book{}
	book.Publisher = publisher

	return book, nil
}

// GetBookByID fetches a Book based on a given ID
func GetBookByID(id string) (model.Book, error) {
	book, err := repository.GetBookByID(id)

	if err != nil {
		return model.Book{}, err
	}

	publisher, _ := repository.GetPublisherByID(strconv.Itoa(int(*book.PublisherID)))
	publisher.Books = []*model.Book{}
	book.Publisher = publisher

	return book, nil
}

// AddBook adds a new Book entity
func AddBook(book model.Book) (model.Book, error) {
	if book.ISBN10 == "" && book.ISBN13 == "" {
		return model.Book{}, errors.New("Must have ISBN 10 and/or ISBN 13.")
	}
	result, err := repository.AddBook(book)

	if err != nil {
		return model.Book{}, err
	}

	publisher, _ := repository.GetPublisherByID(strconv.Itoa(int(*book.PublisherID)))
	result.Publisher = publisher
	mq.PublishBook(mq.BookQueue, result)

	return result, nil
}

// EditBook edits a Book entity
func EditBook(book model.Book, id string) (model.Book, error) {
	book, err := repository.EditBook(book, id)

	if err != nil {
		return model.Book{}, err
	}

	return book, nil
}

// DeleteBook deletes a Book entity
func DeleteBook(id string) error {
	err := repository.DeleteBook(id)

	if err != nil {
		return err
	}

	return nil
}
