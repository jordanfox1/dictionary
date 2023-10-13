package storage_interaction

import (
	"dictionary/dictionary/models"
	"testing"
)

// This just tests that the storage interaction layer correctly calls the storage layer functions.
// We can not test the actual implementations of the storage layer without actually saving something and asserting it has been saved, this does not do that.
func TestSaveDefInDB(t *testing.T) {
	mockStorage := &MockStorage{}

	def := models.CustomDefinition{
		Word:     "test",
		Phonetic: "/test/",
		Meanings: []models.Meaning{},
	}

	err := SaveDefInDB(mockStorage, def)

	if mockStorage.SavedDefinition.Word != def.Word {
		t.Errorf("SaveDefInDB did not store the definition's Word correctly")
	}
	if mockStorage.SavedDefinition.Phonetic != def.Phonetic {
		t.Errorf("SaveDefInDB did not store the definition's Phonetic correctly")
	}

	if err != nil {
		t.Errorf("SaveDefInDB returned an error: %v", err)
	}
}
