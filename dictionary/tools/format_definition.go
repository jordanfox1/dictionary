package tools

import (
	"encoding/json"
	"fmt"
	"log"
)

type Definition struct {
	Definition string   `json:"definition"`
	Synonyms   []string `json:"synonyms"`
	Antonyms   []string `json:"antonyms"`
}

type Meaning struct {
	PartOfSpeech string `json:"partOfSpeech"`
	Definitions  []Definition
	Synonyms     []string `json:"synonyms"`
	Antonyms     []string `json:"antonyms"`
}

type Phonetic struct {
	Text  string `json:"text"`
	Audio string `json:"audio"`
}

type CustomDefinition struct {
	Word     string    `json:"word"`
	Phonetic string    `json:"phonetic"`
	Meanings []Meaning `json:"meanings"`
}

func FormatDefinition(jsonData string) CustomDefinition {
	var customDef []CustomDefinition

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
