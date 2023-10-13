package mocks

import "dictionary/dictionary/models"

var HelmetOutput = models.CustomDefinition{
	Word:     "helmet",
	Phonetic: "/ˈhɛlmət/",
	Meanings: []models.Meaning{
		{
			PartOfSpeech: "noun",
			Synonyms:     []string{"brain bucket", "hard hat"},
			Antonyms:     []string{},
			Definitions: []models.Definition{
				{
					Definition: "A protective head covering, usually part of armour.",
					Synonyms:   []string{},
					Antonyms:   []string{},
				},
				{
					Definition: "That which resembles a helmet in form, position, etc.",
					Synonyms:   []string{},
					Antonyms:   []string{},
				},
			},
		},
		{
			PartOfSpeech: "verb",
			Synonyms:     []string{},
			Antonyms:     []string{},
			Definitions: []models.Definition{
				{
					Definition: "To cover with, or as if with, a helmet.",
					Synonyms:   []string{},
					Antonyms:   []string{},
				},
			},
		},
	},
}
