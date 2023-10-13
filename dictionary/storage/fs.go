package storage

import (
	"dictionary/dictionary/models"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

// holds the functions for interacting with local fs
// FileStorage implements the Storage interface for local file storage
type FileStorage struct {
}

func NewFileStorage() *FileStorage {
	return &FileStorage{}
}

func (fs *FileStorage) Save(def models.CustomDefinition) error {
	// Marshal the CustomDefinition into a JSON string
	jsonData, err := json.Marshal(def)
	if err != nil {
		log.Fatal(err)
		return err
	}

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Construct the absolute path for saving the file in the current directory
	filename := filepath.Join(currentDir+"/saved_data/", def.Word+".json")

	// Create or open the file for writing
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()

	// Write the JSON data to the file
	_, err = file.Write(jsonData)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (fs *FileStorage) Get(word string) (models.CustomDefinition, error) {
	// Implement the file retrieval logic
	return models.CustomDefinition{}, nil
}

func (fs *FileStorage) Delete(filename string) error {
	// Implement the file deletion logic
	return nil

}
