package repository

import (
	"github.com/kirkalyn13/xyz-books-app/pkg/db"
	"github.com/kirkalyn13/xyz-books-app/pkg/model"
)

// GetAuthors fetches the list of Authors from the database
func GetAuthors(searchQuery string) ([]model.Author, error) {
	var authors []model.Author

	if result := db.DB.Preload("Books").
		Where("first_name LIKE ? OR last_name LIKE ? OR middle_name LIKE ?", "%"+searchQuery+"%", "%"+searchQuery+"%", "%"+searchQuery+"%").
		Find(&authors); result.Error != nil {
		return []model.Author{}, result.Error
	}

	return authors, nil
}

// GetAuthorsByID fetches a Author based on a given ID from the database
func GetAuthorByID(id string) (model.Author, error) {
	var author model.Author

	if result := db.DB.Preload("Books").First(&author, id); result.Error != nil {
		return model.Author{}, result.Error
	}

	return author, nil
}

// AddAuthor adds a new Author entity from the database
func AddAuthor(author model.Author) (model.Author, error) {
	if result := db.DB.Create(&author); result.Error != nil {
		return model.Author{}, result.Error
	}

	return author, nil
}

// EditAuthor edits a Author entity from the database
func EditAuthor(author model.Author, id string) (model.Author, error) {
	if result := db.DB.Where("id = ?", id).Updates(&author); result.Error != nil {
		return model.Author{}, result.Error
	}

	return author, nil
}

// DeleteAuthor deletes a Author entity from the database
func DeleteAuthor(id string) error {
	if result := db.DB.Delete(&model.Author{}, "id = ?", id); result.Error != nil {
		return result.Error
	}

	return nil
}
