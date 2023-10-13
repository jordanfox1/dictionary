package main

import (
	"dictionary/dictionary/models"
	"dictionary/dictionary/network"
	"dictionary/dictionary/network_interaction"
	"dictionary/dictionary/storage"
	"dictionary/dictionary/storage_interaction"
	"dictionary/dictionary/tools"
	"dictionary/dictionary/user_input"
	"fmt"
	"log"
)

var fileStorage = storage.NewFileStorage()

func main() {
	user_input.PromptUser(user_input.WelcomePrompt)
	input := user_input.GetUserInput("Enter a word: ")
	validInput, err := user_input.ValidateUserInput(input)
	if err != nil {
		log.Fatal(err)
	}

	// Create an instance of the real network client to use it's implementation.
	client := &network.RealNetworkClient{}

	// Create an instance of the NetworkInteraction with the real client.
	ni := network_interaction.NetworkInteraction{Client: client}
	rawDefinition, err := ni.GetDefinition(validInput, "https://api.dictionaryapi.dev/api/v2/entries/en/")
	if err != nil {
		log.Fatal(err)
	}

	def := tools.FormatDefinition(rawDefinition)
	gottenDef, err := getSavedDef(def.Word)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gottenDef)
}

func saveDef(def models.CustomDefinition) {
	user_input.PromptUser(user_input.EnterOperationPrompt)
	storage_interaction.SaveDefInDB(fileStorage, def)
}

func getSavedDef(word string) (models.CustomDefinition, error) {
	return fileStorage.Get(word)
}
