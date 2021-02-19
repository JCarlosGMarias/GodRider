package webclients

import (
	"fmt"
	"godrider/webclients/webclientmodels"
)

// WebClientFactorier defines a factory struct which gives us adapted webclients for our consumption
type WebClientFactorier interface {
	// GetClient should return a WebClient implementation based on our ClientData params
	GetClient(clientData *webclientmodels.ClientData) (WebClient, error)
}

// WebClientFactory is WebClientFactorier's implementation struct
type WebClientFactory struct {
	webClients map[int]WebClient
}

// ClientFactory is WebClientFactorier's implementation instance
var ClientFactory = WebClientFactory{
	webClients: map[int]WebClient{
		5: &mockWebClient{
			mockUser:     "Fooscott",
			mockPassword: "Bartiger",
		},
	},
}

// GetClient returns a WebClient implementation based on our ClientData params
func (builder *WebClientFactory) GetClient(clientData *webclientmodels.ClientData) (WebClient, error) {
	client, ok := builder.webClients[clientData.ProviderID]
	if !ok {
		return nil, fmt.Errorf("WebClient with ID %d not registered", clientData.ProviderID)
	}

	success, err := client.Login(clientData)
	if success {
		return client, nil
	}
	return nil, err
}
