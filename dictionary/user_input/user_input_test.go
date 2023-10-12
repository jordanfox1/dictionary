package user_input

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestPromptUserForInput(t *testing.T) {
	// Capture the standard output.
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Restore the original standard output at the end of the test.
	defer func() {
		os.Stdout = oldStdout
	}()

	// Call the function you want to test.
	PromptUserForInput(WelcomePrompt)

	// Close the write pipe to read the captured output.
	w.Close()

	// Read and check the captured output.
	capturedOutput, _ := io.ReadAll(r)

	expectedOutput := "---------- ========== Welcome ========== ----------\n"
	if string(capturedOutput) != expectedOutput {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, capturedOutput)
	}
}

func TestGetUserInput(t *testing.T) {
	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r

	go func() {
		fmt.Fprint(w, "Testing\n")
		w.Close()
	}()

	userInput := GetUserInput()

	os.Stdin = oldStdin

	if userInput != "Testing" {
		t.Errorf("Expected 'Test Input', but got '%s'", userInput)
	}
}
