package repository

import (
	"github.com/kirkalyn13/xyz-books-app/pkg/db"
	"github.com/kirkalyn13/xyz-books-app/pkg/model"
)

// GetAuthors fetches the list of Authors from the database
func GetAuthors() ([]model.Author, error) {
	var authors []model.Author

	result := db.DB.Model(&model.Author{}).Preload("Books").Find(&authors)

	if result.Error != nil {
		return []model.Author{}, result.Error
	}

	return authors, nil
}

// GetAuthorsByID fetches a Author based on a given ID from the database
func GetAuthorByID(id string) (model.Author, error) {
	var author model.Author

	result := db.DB.Model(&model.Author{}).Preload("Books").Find(&author, id)

	if result.Error != nil {
		return model.Author{}, result.Error
	}

	return author, nil
}

// AddAuthor adds a new Author entity from the database
func AddAuthor(author model.Author) (model.Author, error) {
	result := db.DB.Unscoped().Create(&author)

	if result.Error != nil {
		return model.Author{}, result.Error
	}

	return author, nil
}

// EditAuthor edits a Author entity from the database
func EditAuthor(author model.Author, id string) (model.Author, error) {
	result := db.DB.First(&author, id)

	if result.Error != nil {
		return model.Author{}, result.Error
	}

	return author, nil
}

// DeleteAuthor deletes a Author entity from the database
func DeleteAuthor(id string) error {
	result := db.DB.Where("id = ?", id).Delete(&model.Author{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
