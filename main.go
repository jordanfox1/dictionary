package main

import (
	"dictionary/dictionary/models"
	"dictionary/dictionary/network"
	"dictionary/dictionary/network_interaction"
	"dictionary/dictionary/storage"
	"dictionary/dictionary/storage_interaction"
	"dictionary/dictionary/tools"
	"dictionary/dictionary/user_input"
	"log"
)

var fileStorage = storage.NewFileStorage()

func main() {
	user_input.PromptUser(user_input.WelcomePrompt)
	input := user_input.GetUserInput()
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
	saveDef(def)
}

func saveDef(def models.CustomDefinition) {
	user_input.PromptUser("")
	storage_interaction.SaveDefInDB(fileStorage, def)
}

func getSavedDef(word string) {

}
