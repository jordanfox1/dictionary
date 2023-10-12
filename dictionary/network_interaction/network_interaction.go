package network_interaction

import "dictionary/dictionary/network" // Import the network package.

type NetworkInteraction struct {
	Client network.NetworkClient
}

func (ni *NetworkInteraction) GetDefinition(word, url string) (string, error) {
	// Use the network client to get the definition.
	return ni.Client.GetDefinition(word, url)
}
