package user_input

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

var (
	Equals10      = strings.Repeat("=", 10)
	Dash10        = strings.Repeat("-", 10)
	WelcomePrompt = fmt.Sprintf("%s %s Welcome %s %s", Dash10, Equals10, Equals10, Dash10)
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

func GetUserInput() string {
	fmt.Print("Enter a word: ")
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
