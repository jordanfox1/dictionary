package tools

import (
	"dictionary/dictionary/mocks"
	"dictionary/dictionary/models"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestFormatDefinition(t *testing.T) {

	fp1 := filepath.Join("../mocks/", "extrinsic.json")
	jsonData1, err := os.ReadFile(fp1)
	if err != nil {
		t.Fatalf("Failed to read JSON file: %v", err)
	}

	fp2 := filepath.Join("../mocks/", "helmet.json")
	jsonData2, err := os.ReadFile(fp2)
	if err != nil {
		t.Fatalf("Failed to read JSON file: %v", err)
	}

	type args struct {
		jsonData string
	}
	tests := []struct {
		name string
		args args
		want models.CustomDefinition
	}{
		{
			name: "case 1",
			args: args{
				jsonData: string(jsonData1),
			},
			want: mocks.ExtrinsicOutput,
		},
		{
			name: "case 2",
			args: args{
				jsonData: string(jsonData2),
			},
			want: mocks.HelmetOutput,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("Asserting got and want are deeply equal")
			if got := FormatDefinition(tt.args.jsonData); !reflect.DeepEqual(got, tt.want) {
				fmt.Printf("got: %+v\n", got)
				fmt.Printf("want: %+v\n", tt.want)
				t.Errorf("FormatDefinition() = %v, want %v", got, tt.want)
			}
		})
	}
}
