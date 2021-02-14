package services

import (
	"godrider/dtos/requests"
	"godrider/dtos/responses"
	"godrider/infrastructures"
	"godrider/infrastructures/models"
)

// ProviderServicer is the provider's service layer
type ProviderServicer interface {
	// GetAllProviders should return all provider's models in a ProviderResponse format
	GetAllProviders() []responses.ProviderResponse
	// GetProviderByID should return an unique provider model by its ID
	GetProviderByID(request *requests.ProviderRequest) (responses.ProviderResponse, error)
}

// ProviderService is ProviderServicer's implementation struct
type ProviderService struct {
	providerInfrastructure infrastructures.ProviderInfrastructurer
}

// ProviderSrv is ProviderServicer's implementation instance
var ProviderSrv ProviderServicer = &ProviderService{providerInfrastructure: infrastructures.ProviderDb}

// GetAllProviders returns all provider's models in a ProviderResponse format
func (service *ProviderService) GetAllProviders() []responses.ProviderResponse {
	providers, count, _ := service.providerInfrastructure.GetAllProviders()

	providersResponses := make([]responses.ProviderResponse, count)
	for _, provider := range providers {
		providersResponse := parseProviderToProviderResponse(&provider)
		providersResponses = append(providersResponses, providersResponse)
	}

	return providersResponses
}

// GetProviderByID returns an unique provider model by its ID
func (service *ProviderService) GetProviderByID(request *requests.ProviderRequest) (responses.ProviderResponse, error) {
	provider, err := service.providerInfrastructure.GetSingleProviderByID(request.ProviderID)
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
