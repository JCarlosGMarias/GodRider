package services

import (
	"godrider/dtos/requests"
	"godrider/dtos/responses"
	"godrider/infrastructures"
	"godrider/infrastructures/models"
)

// ProviderServicer is the provider's service layer
type ProviderServicer interface {
	// GetAll should return all provider's models in a ProviderResponse format
	GetAll() []responses.ProviderResponse
	// GetByID should return an unique provider model by its ID
	GetByID(request *requests.ProviderRequest) (responses.ProviderResponse, error)
}

// ProviderService is ProviderServicer's implementation struct
type ProviderService struct {
	providerInfrastructure infrastructures.ProviderInfrastructurer
}

// GetAll returns all provider's models in a ProviderResponse format
func (s *ProviderService) GetAll() []responses.ProviderResponse {
	providers, count, _ := s.providerInfrastructure.GetAll()

	providersRs := make([]responses.ProviderResponse, count)
	for i, provider := range providers {
		providerRs := parseProviderToProviderResponse(&provider)
		providersRs[i] = providerRs
	}

	return providersRs
}

// GetByID returns an unique provider model by its ID
func (s *ProviderService) GetByID(request *requests.ProviderRequest) (responses.ProviderResponse, error) {
	provider, err := s.providerInfrastructure.GetSingleByID(request.ProviderID)
	if err != nil || provider.ID == 0 {
		return responses.ProviderResponse{}, &responses.ErrorResponse{Code: responses.REGISTER_NOT_FOUND, Message: "Provider not found!"}
	}

	providerRs := parseProviderToProviderResponse(&provider)
	return providerRs, nil
}

// ProviderInfrastructure setter
func (s *ProviderService) ProviderInfrastructure(istruct *infrastructures.ProviderInfrastructurer) {
	if s.providerInfrastructure == nil {
		s.providerInfrastructure = *istruct
	}
}

func parseProviderToProviderResponse(p *models.Provider) responses.ProviderResponse {
	return responses.ProviderResponse{
		ID:      p.ID,
		Name:    p.Name,
		Contact: p.Contact,
	}
}
