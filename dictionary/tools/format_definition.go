package tools

import (
	"dictionary/dictionary/models"
	"encoding/json"
	"fmt"
	"log"
)

func FormatDefinition(jsonData string) models.CustomDefinition {
	var customDef []models.CustomDefinition

	if err := json.Unmarshal([]byte(jsonData), &customDef); err != nil {
		log.Fatalf("JSON parsing error: %v", err)
	}

	if len(customDef) > 0 {
		custom := customDef[0]
		fmt.Printf("Word: %s\n", custom.Word)
		fmt.Printf("Phonetic: %s\n", custom.Phonetic)

		for _, meaning := range custom.Meanings {
			fmt.Printf("Part of Speech: %s\n", meaning.PartOfSpeech)
			fmt.Printf("Synonyms: %v\n", meaning.Synonyms)
			fmt.Printf("Antonyms: %v\n", meaning.Antonyms)
			for _, def := range meaning.Definitions {
				fmt.Printf("Definition: %s\n", def.Definition)
				fmt.Printf("Synonyms: %v\n", def.Synonyms)
				fmt.Printf("Antonyms: %v\n", def.Antonyms)
			}
		}
	}

	return customDef[0]
}
