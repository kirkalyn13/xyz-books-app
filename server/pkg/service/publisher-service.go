package service

import (
	"github.com/kirkalyn13/xyz-books-app/pkg/model"
	"github.com/kirkalyn13/xyz-books-app/pkg/repository"
)

// GetPublishers fetches the list of Publishers
func GetPublishers() []model.Publisher {
	return repository.GetPublishers()
}

// GetPublishersByID fetches a Publisher based on a given ID
func GetPublisherByID(id string) model.Publisher {
	return repository.GetPublisherByID(id)
}

// AddPublisher adds a new Publisher entity
func AddPublisher(Publisher model.Publisher) (model.Publisher, error) {
	result, err := repository.AddPublisher(Publisher)

	if err != nil {
		return model.Publisher{}, err
	}

	return result, nil
}

// EditPublisher edits a Publisher entity
func EditPublisher(Publisher model.Publisher, id string) model.Publisher {
	result := repository.EditPublisher(Publisher, id)

	return result
}

// DeletePublisher deletes a Publisher entity
func DeletePublisher(id string) {
	repository.DeletePublisher(id)
}
