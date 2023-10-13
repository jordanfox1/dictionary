package storage_interaction

import (
	"dictionary/dictionary/models"
	"errors"
)

type MockStorage struct {
	SavedDefinition models.CustomDefinition
	// You can add fields for additional testing if needed
}

func (ms *MockStorage) Save(def models.CustomDefinition) error {
	ms.SavedDefinition = def
	// Mock the Save method behavior
	return nil
}

func (ms *MockStorage) Get(word string) (models.CustomDefinition, error) {
	// Implement the Get method behavior
	if ms.SavedDefinition.Word != word {
		return ms.SavedDefinition, errors.New("no definition for word was found")
	}
	return ms.SavedDefinition, nil
}

func (ms *MockStorage) Delete(word string) error {
	// Implement the Delete method behavior
	return nil
}
