package main

import "dictionary/dictionary/user_input"

func main() {
	user_input.PromptUser(user_input.WelcomePrompt)
	input := user_input.GetUserInput()
	user_input.ValidateUserInput(input)
}
