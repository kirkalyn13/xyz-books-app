package repository

import (
	"github.com/kirkalyn13/xyz-books-app/pkg/db"
	"github.com/kirkalyn13/xyz-books-app/pkg/model"
)

// GetPublishers fetches the list of Publishers from the database
func GetPublishers() ([]model.Publisher, error) {
	var publishers []model.Publisher

	result := db.DB.Preload("Books").Find(&publishers)

	if result.Error != nil {
		return []model.Publisher{}, result.Error
	}

	return publishers, nil
}

// GetPublishersByID fetches a Publisher based on a given ID from the database
func GetPublisherByID(id string) (model.Publisher, error) {
	var publisher model.Publisher

	result := db.DB.Preload("Books").Find(&publisher, id)

	if result.Error != nil {
		return model.Publisher{}, result.Error
	}

	return publisher, nil
}

// AddPublisher adds a new Publisher entity from the database
func AddPublisher(publisher model.Publisher) (model.Publisher, error) {
	result := db.DB.Unscoped().Create(&publisher)

	if result.Error != nil {
		return model.Publisher{}, result.Error
	}

	return publisher, nil
}

// EditPublisher edits a Publisher entity from the database
func EditPublisher(publisher model.Publisher, id string) (model.Publisher, error) {
	result := db.DB.First(&publisher, id)

	if result.Error != nil {
		return model.Publisher{}, result.Error
	}

	return publisher, nil
}

// DeletePublisher deletes a Publisher entity from the database
func DeletePublisher(id string) error {
	result := db.DB.Where("id = ?", id).Delete(&model.Publisher{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
