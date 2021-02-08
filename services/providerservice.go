package services

import (
	"godrider/dtos/requests"
	"godrider/dtos/responses"
	"godrider/infrastructures"
	"godrider/infrastructures/models"
)

type ProviderService struct {
	providerInfrastructure infrastructures.ProvidersInfrastructure
}

var ProviderSrv = ProviderService{
	providerInfrastructure: infrastructures.ProvidersDb,
}

func (service *ProviderService) GetAllProviders() []responses.ProviderResponse {
	providers, count, _ := service.providerInfrastructure.GetAllProviders()

	providersResponses := make([]responses.ProviderResponse, count)
	for _, provider := range providers {
		providersResponse := parseProviderToProviderResponse(&provider)
		providersResponses = append(providersResponses, providersResponse)
	}

	return providersResponses
}

func (service *ProviderService) GetProviderById(request requests.ProviderRequest) (responses.ProviderResponse, error) {
	provider, err := service.providerInfrastructure.GetSingleProviderById(request.ProviderID)
	if err != nil || provider.ID == 0 {
		return responses.ProviderResponse{}, &responses.ErrorResponse{Code: responses.REGISTER_NOT_FOUND, Message: "Provider not found!"}
	}

	providerResponse := parseProviderToProviderResponse(&provider)
	return providerResponse, nil
}

func parseProviderToProviderResponse(provider *models.Provider) responses.ProviderResponse {
	return responses.ProviderResponse{
		ID:      provider.ID,
		Name:    provider.Name,
		Contact: provider.Contact,
	}
}
