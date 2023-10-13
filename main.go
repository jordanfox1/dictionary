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

	inputOp := user_input.GetUserInput(user_input.EnterOperationPrompt)
	validInputOp, err := user_input.ValidateUserOpInput(inputOp)
	if err != nil {
		log.Fatal(err)
	}

	inputWord := user_input.GetUserInput("Enter a word: ")
	validInputWord, err := user_input.ValidateUserInput(inputWord)
	if err != nil {
		log.Fatal(err)
	}

	if validInputOp == "l" {
		getNewDef(validInputWord)
	}

	if validInputOp == "g" {
		gottenDef, err := getSavedDef(validInputWord)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Got definition:")
		fmt.Println(gottenDef)
	}

}

func getNewDef(word string) models.CustomDefinition {
	// Create an instance of the real network client to use it's implementation.
	client := &network.RealNetworkClient{}

	// Create an instance of the NetworkInteraction with the real client.
	ni := network_interaction.NetworkInteraction{Client: client}
	rawDefinition, err := ni.GetDefinition(word, "https://api.dictionaryapi.dev/api/v2/entries/en/")
	if err != nil {
		log.Fatal(err)
	}

	def := tools.FormatDefinition(rawDefinition)
	err = storage_interaction.SaveDefInDB(fileStorage, def)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Definition for %s has been saved!", word)
	return def
}

func getSavedDef(word string) (models.CustomDefinition, error) {
	return fileStorage.Get(word)
}
