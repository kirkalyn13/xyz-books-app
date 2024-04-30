package repository

import (
	"github.com/kirkalyn13/xyz-books-app/pkg/db"
	"github.com/kirkalyn13/xyz-books-app/pkg/model"
)

// GetAuthors fetches the list of Authors from the database
func GetAuthors() []model.Author {
	var Authors []model.Author

	db.DB.Unscoped().Find(&Authors)

	return Authors
}

// GetAuthorsByID fetches a Author based on a given ID from the database
func GetAuthorByID(id string) model.Author {
	var Author model.Author

	db.DB.Unscoped().Find(&Author, id)

	return Author
}

// AddAuthor adds a new Author entity from the database
func AddAuthor(Author model.Author) (model.Author, error) {
	result := db.DB.Unscoped().Create(&Author)

	if result.Error != nil {
		return model.Author{}, result.Error
	}

	return Author, nil
}

// EditAuthor edits a Author entity from the database
func EditAuthor(Author model.Author, id string) model.Author {
	db.DB.First(&Author, id)

	return Author
}

// DeleteAuthor deletes a Author entity from the database
func DeleteAuthor(id string) {
	db.DB.Delete(&model.Author{}, id)
}
