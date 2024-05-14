package repository

import (
	"github.com/kirkalyn13/xyz-books-app/server/pkg/db"
	"github.com/kirkalyn13/xyz-books-app/server/pkg/model"
)

// GetPublishers fetches the list of Publishers from the database
func GetPublishers(searchQuery string) ([]model.Publisher, error) {
	var publishers []model.Publisher

	if result := db.DB.Preload("Books").Where("name LIKE ?", "%"+searchQuery+"%").Find(&publishers); result.Error != nil {
		return []model.Publisher{}, result.Error
	}

	return publishers, nil
}

// GetPublisherByID fetches a Publisher based on a given ID from the database
func GetPublisherByID(id string) (model.Publisher, error) {
	var publisher model.Publisher

	if result := db.DB.Preload("Books").First(&publisher, id); result.Error != nil {
		return model.Publisher{}, result.Error
	}

	return publisher, nil
}

// AddPublisher adds a new Publisher entity from the database
func AddPublisher(publisher model.Publisher) (model.Publisher, error) {
	if result := db.DB.Create(&publisher); result.Error != nil {
		return model.Publisher{}, result.Error
	}

	return publisher, nil
}

// EditPublisher edits a Publisher entity from the database
func EditPublisher(publisher model.Publisher, id string) (model.Publisher, error) {
	if result := db.DB.Where("id = ?", id).First(&model.Publisher{}); result.Error != nil {
		return model.Publisher{}, result.Error
	}

	if result := db.DB.Where("id = ?", id).Updates(&publisher); result.Error != nil {
		return model.Publisher{}, result.Error
	}

	return publisher, nil
}

// DeletePublisher deletes a Publisher entity from the database
func DeletePublisher(id string) error {
	if result := db.DB.Where("id = ?", id).First(&model.Publisher{}); result.Error != nil {
		return result.Error
	}

	if result := db.DB.Delete(&model.Publisher{}, "id = ?", id); result.Error != nil {
		return result.Error
	}

	return nil
}
