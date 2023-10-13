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
	PromptUser(WelcomePrompt)

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

	userInput := GetUserInput("")

	os.Stdin = oldStdin

	if userInput != "Testing" {
		t.Errorf("Expected 'Test Input', but got '%s'", userInput)
	}
}

func TestValidateUserInput(t *testing.T) {
	type args struct {
		input       string
		expectedErr error
		expected    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				input:       "",
				expectedErr: ErrNoInput,
				expected:    "",
			},
			name:    "Test Case 1",
			wantErr: true,
		},
		{
			args: args{
				input:       "0",
				expectedErr: ErrNoNumeric,
				expected:    "",
			},
			name:    "Test Case 2",
			wantErr: true,
		},
		{
			args: args{
				input:       "big shot",
				expectedErr: ErrNoNumeric,
				expected:    "big shot",
			},
			name:    "Test Case 3",
			wantErr: false,
		},
		{
			args: args{
				input:       "HEY",
				expectedErr: ErrNoNumeric,
				expected:    "HEY",
			},
			name:    "Test Case 4",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			actual, err := ValidateUserInput(tt.args.input)
			if err != nil && tt.wantErr {
				if err != tt.args.expectedErr {
					t.Errorf("ValidateUserInput() error = %v, wantErr %v, expected error message to be %v", err, tt.wantErr, tt.args.expectedErr)
				}
			}

			if actual != tt.args.expected {
				t.Errorf("ValidateUserInput() error = %v, wantErr %v, expected output to be %v", err, tt.wantErr, tt.args.expected)
			}

		})
	}
}

func TestValidateUserOpInput(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				input: "",
			},
			wantErr: true,
		},
		{
			name: "test case 2",
			args: args{
				input: "W",
			},
			wantErr: true,
		},
		{
			name: "test case 3",
			args: args{
				input: "l",
			},
			wantErr: false,
			want:    "l",
		},
		{
			name: "test case 4",
			args: args{
				input: "g",
			},
			wantErr: false,
			want:    "g",
		},
		{
			name: "test case 5",
			args: args{
				input: "L",
			},
			wantErr: false,
			want:    "l",
		},
		{
			name: "test case 6",
			args: args{
				input: "G",
			},
			wantErr: false,
			want:    "g",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateUserOpInput(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateUserOpInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ValidateUserOpInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
