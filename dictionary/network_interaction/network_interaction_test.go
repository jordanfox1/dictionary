package network_interaction

import "testing"

type MockNetworkClient struct {
	Definition string
	Err        error
}

func (c *MockNetworkClient) GetDefinition(word, url string) (string, error) {
	return c.Definition, c.Err
}

func TestGetDefinition(t *testing.T) {

	mockClient := &MockNetworkClient{
		Definition: "a greeting",
		Err:        nil,
	}

	ni := NetworkInteraction{Client: mockClient}
	definition, err := ni.GetDefinition("test input", "test url")

	if err != nil || definition != mockClient.Definition {
		t.Fatal(err)
	}

}
