package service

import (
	"github.com/kirkalyn13/xyz-books-app/pkg/model"
	"github.com/kirkalyn13/xyz-books-app/pkg/repository"
)

// GetAuthors fetches the list of Authors
func GetAuthors(searchQuery string) ([]model.Author, error) {
	authors, err := repository.GetAuthors(searchQuery)

	if err != nil {
		return []model.Author{}, err
	}

	return authors, err
}

// GetAuthorsByID fetches a Author based on a given ID
func GetAuthorByID(id string) (model.Author, error) {
	author, err := repository.GetAuthorByID(id)

	if err != nil {
		return model.Author{}, err
	}

	return author, nil
}

// AddAuthor adds a new Author entity
func AddAuthor(author model.Author) (model.Author, error) {
	result, err := repository.AddAuthor(author)

	if err != nil {
		return model.Author{}, err
	}

	return result, nil
}

// EditAuthor edits a Author entity
func EditAuthor(author model.Author, id string) (model.Author, error) {
	result, err := repository.EditAuthor(author, id)

	if err != nil {
		return model.Author{}, err
	}

	return result, nil
}

// DeleteAuthor deletes a Author entity
func DeleteAuthor(id string) error {
	err := repository.DeleteAuthor(id)

	if err != nil {
		return err
	}

	return nil
}
