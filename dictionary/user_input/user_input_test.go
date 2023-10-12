package user_input

import (
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
