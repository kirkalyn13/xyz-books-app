package repository

import (
	"github.com/kirkalyn13/xyz-books-app/pkg/db"
	"github.com/kirkalyn13/xyz-books-app/pkg/model"
)

// GetPublishers fetches the list of Publishers from the database
func GetPublishers() []model.Publisher {
	var Publishers []model.Publisher

	db.DB.Unscoped().Find(&Publishers)

	return Publishers
}

// GetPublishersByID fetches a Publisher based on a given ID from the database
func GetPublisherByID(id string) model.Publisher {
	var Publisher model.Publisher

	db.DB.Unscoped().Find(&Publisher, id)

	return Publisher
}

// AddPublisher adds a new Publisher entity from the database
func AddPublisher(Publisher model.Publisher) (model.Publisher, error) {
	result := db.DB.Unscoped().Create(&Publisher)

	if result.Error != nil {
		return model.Publisher{}, result.Error
	}

	return Publisher, nil
}

// EditPublisher edits a Publisher entity from the database
func EditPublisher(Publisher model.Publisher, id string) model.Publisher {
	db.DB.First(&Publisher, id)

	return Publisher
}

// DeletePublisher deletes a Publisher entity from the database
func DeletePublisher(id string) {
	db.DB.Delete(&model.Publisher{}, id)
}
