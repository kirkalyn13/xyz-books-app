package service

import (
	"github.com/kirkalyn13/xyz-books-app/pkg/model"
	"github.com/kirkalyn13/xyz-books-app/pkg/repository"
)

// GetPublishers fetches the list of Publishers
func GetPublishers() ([]model.Publisher, error) {
	publishers, err := repository.GetPublishers()

	if err != nil {
		return []model.Publisher{}, err
	}

	return publishers, nil
}

// GetPublishersByID fetches a Publisher based on a given ID
func GetPublisherByID(id string) (model.Publisher, error) {
	publisher, err := repository.GetPublisherByID(id)

	if err != nil {
		return model.Publisher{}, err
	}

	return publisher, nil
}

// AddPublisher adds a new Publisher entity
func AddPublisher(publisher model.Publisher) (model.Publisher, error) {
	result, err := repository.AddPublisher(publisher)

	if err != nil {
		return model.Publisher{}, err
	}

	return result, nil
}

// EditPublisher edits a Publisher entity
func EditPublisher(publisher model.Publisher, id string) (model.Publisher, error) {
	result, err := repository.EditPublisher(publisher, id)

	if err != nil {
		return model.Publisher{}, err
	}

	return result, nil
}

// DeletePublisher deletes a Publisher entity
func DeletePublisher(id string) error {
	err := repository.DeletePublisher(id)

	if err != nil {
		return err
	}

	return nil
}
