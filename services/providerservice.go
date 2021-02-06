package services

import (
	"godrider/dtos/requests"
	"godrider/dtos/responses"
	"godrider/infrastructures/models"
)

type ProviderService struct {
	providerDb []models.Provider
}

var ProviderSrv = ProviderService{
	providerDb: []models.Provider{
		{ID: 1, Name: "Balloon", Contact: "ballooncorp@evil.death"},
		{ID: 2, Name: "PanzerChomps", Contact: "panzercorp@evil.death"},
		{ID: 3, Name: "SimplyDelight", Contact: "simplydelightcorp@evil.death"},
		{ID: 4, Name: "Nom", Contact: "nomcorp@evil.death"},
	},
}

func (service *ProviderService) GetAllProviders() []models.Provider {
	return service.providerDb
}

func (service *ProviderService) GetProviderById(request requests.ProviderRequest) (models.Provider, responses.ErrorResponse) {
	for _, provider := range service.providerDb {
		if provider.ID == request.ProviderID {
			return provider, responses.ErrorResponse{}
		}
	}
	return models.Provider{}, responses.ErrorResponse{Code: 4, Message: "Provider not found!"}
}
