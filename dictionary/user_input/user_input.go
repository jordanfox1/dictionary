package user_input

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

var (
	Equals10             = strings.Repeat("=", 10)
	Dash10               = strings.Repeat("-", 10)
	WelcomePrompt        = fmt.Sprintf("%s %s Welcome %s %s", Dash10, Equals10, Equals10, Dash10)
	EnterOperationPrompt = "What would you like to do (l for lookup new word, g for get a saved word)"

	SavePrompt   = "Would you like to save this definition? (y for yes, n no)"
	SavingPrompt = fmt.Sprintf("%s %s Saving definition %s %s", Dash10, Equals10, Equals10, Dash10)
	SavedPrompt  = "Definition has been saved!"

	ErrValidation = errors.New("error validating input")
	ErrNoInput    = errors.New("no input detected")
	ErrNoNumeric  = errors.New("please enter only letters")
)

func PromptUser(prompt string) {
	fmt.Println(prompt)
}

type InputReader interface {
	Read(p []byte) (n int, err error)
}

func GetUserInput(prompt string) string {
	fmt.Print(prompt)
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Error reading input:", err)
	}
	return input
}

func ValidateUserInput(input string) (string, error) {
	if input == "" {
		return "", ErrNoInput
	}

	for _, r := range input {
		if !unicode.IsLetter(r) && !unicode.IsSpace(r) && r != '-' {
			fmt.Printf("Error: Rune '%c' is not a letter, whitespace, or '-'.\n", r)
			return "", ErrNoNumeric
		}
	}

	fmt.Printf("attempting to get a definition for %s\n", input)
	return input, nil
}

func ValidateUserOpInput(input string) (string, error) {
	switch input {
	case "l":
		return input, nil
	case "g":
		return input, nil
	case "L":
		return "l", nil
	case "G":
		return "g", nil
	default:
		return "", fmt.Errorf("(l for lookup new word, g for get a saved word), you entered %s", input)
	}
}
