// In dictionary/dictionary/network/network.go

package network

type NetworkClient interface {
	GetDefinition(word, url string) (string, error)
}

type RealNetworkClient struct {
	// You can add any configuration or dependencies needed.
}

func (c *RealNetworkClient) GetDefinition(word, url string) (string, error) {
	// Make the real HTTP request here and return the result.
	return "", nil
}
