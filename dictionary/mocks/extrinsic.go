package mocks

import (
	"dictionary/dictionary/models"
)

var ExtrinsicOutput = models.CustomDefinition{
	Word:     "extrinsic",
	Phonetic: "/ɛks.ˈtɹɪn.zɪk/",
	Meanings: []models.Meaning{
		{
			PartOfSpeech: "noun",
			Definitions: []models.Definition{
				{
					Definition: "An external factor",
					Synonyms:   []string{},
					Antonyms:   []string{},
				},
			},
			Synonyms: []string{},
			Antonyms: []string{},
		},
		{
			PartOfSpeech: "adjective",
			Definitions: []models.Definition{
				{
					Definition: "External; separable from the thing itself; inessential",
					Synonyms:   []string{},
					Antonyms:   []string{},
				},
				{
					Definition: "Not belonging to something; outside",
					Synonyms:   []string{},
					Antonyms:   []string{},
				},
			},
			Synonyms: []string{},
			Antonyms: []string{"inherent", "intrinsic"},
		},
	},
}
