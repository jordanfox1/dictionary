package storage_interaction

import "dictionary/dictionary/models"

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
	return models.CustomDefinition{}, nil
}

func (ms *MockStorage) Delete(word string) error {
	// Implement the Delete method behavior
	return nil
}
