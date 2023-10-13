// In dictionary/dictionary/network/network.go

package network

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type NetworkClient interface {
	GetDefinition(word, url string) (string, error)
}

type RealNetworkClient struct {
	// You can add any configuration or dependencies needed.
}

func (c *RealNetworkClient) GetDefinition(word, url string) (string, error) {
	client := &http.Client{}

	apiUrl := fmt.Sprintf("%s%s", url, word)

	// Make the HTTP GET request.
	resp, err := client.Get(apiUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read and parse the response.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// The response body contains the definition data in JSON format.
	// You can parse it and extract the definition as needed.
	// For simplicity, let's assume the response body is the definition.
	definition := string(body)

	return definition, nil
}
