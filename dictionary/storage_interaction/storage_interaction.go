package storage_interaction

import "dictionary/dictionary/models"

// Storage interface
type Storage interface {
	Save(def models.CustomDefinition) error
	Get(word string) (models.CustomDefinition, error)
	Delete(word string) error
}

func SaveDefInDB(storage Storage, def models.CustomDefinition) error {
	return storage.Save(def)
}
