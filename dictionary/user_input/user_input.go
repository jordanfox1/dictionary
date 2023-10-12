package user_input

import (
	"fmt"
	"strings"
)

var (
	Equals10      = strings.Repeat("=", 10)
	Dash10        = strings.Repeat("-", 10)
	WelcomePrompt = fmt.Sprintf("%s %s Welcome %s %s", Dash10, Equals10, Equals10, Dash10)
)

func PromptUserForInput(prompt string) {
	fmt.Println(prompt)
}
