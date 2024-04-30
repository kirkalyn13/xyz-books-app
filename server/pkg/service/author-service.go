package service

import (
	"github.com/kirkalyn13/xyz-books-app/pkg/model"
	"github.com/kirkalyn13/xyz-books-app/pkg/repository"
)

// GetAuthors fetches the list of Authors
func GetAuthors() []model.Author {
	return repository.GetAuthors()
}

// GetAuthorsByID fetches a Author based on a given ID
func GetAuthorByID(id string) model.Author {
	return repository.GetAuthorByID(id)
}

// AddAuthor adds a new Author entity
func AddAuthor(Author model.Author) (model.Author, error) {
	result, err := repository.AddAuthor(Author)

	if err != nil {
		return model.Author{}, err
	}

	return result, nil
}

// EditAuthor edits a Author entity
func EditAuthor(Author model.Author, id string) model.Author {
	result := repository.EditAuthor(Author, id)

	return result
}

// DeleteAuthor deletes a Author entity
func DeleteAuthor(id string) {
	repository.DeleteAuthor(id)
}
