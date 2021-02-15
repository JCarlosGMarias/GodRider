package webclients

import (
	"fmt"
	"godrider/webclients/webclientmodels"
)

type WebClientFactory struct {
	webClients map[int]WebClient
}

var ClientFactory = WebClientFactory{
	webClients: map[int]WebClient{
		5: &mock,
	},
}

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
