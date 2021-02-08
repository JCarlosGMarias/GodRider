package services

import (
	"godrider/dtos/requests"
	"godrider/dtos/responses"
	"godrider/infrastructures"
	"godrider/infrastructures/models"
)

type UserProviderService struct {
	usersInfrastructure        infrastructures.UsersInfrastructure
	providerInfrastructure     infrastructures.ProvidersInfrastructure
	userProviderInfrastructure infrastructures.UserProviderInfrastructure
}

var UserProviderSrv = UserProviderService{
	usersInfrastructure:        infrastructures.UsersDb,
	providerInfrastructure:     infrastructures.ProvidersDb,
	userProviderInfrastructure: infrastructures.UserProviderDb,
}

func (service *UserProviderService) AddSuscription(request *requests.UserProviderRequest) error {
	model := models.UserProvider{
		UserId:     request.UserId,
		ProviderId: request.ProviderId,
	}
	if request.IsActive {
		model.IsActive = 1
	} else {
		model.IsActive = 0
	}

	err := service.userProviderInfrastructure.InsertSingle(&model)
	if err == nil {
		return &responses.ErrorResponse{Code: responses.OK, Message: ""}
	}
	return &responses.ErrorResponse{Code: responses.ADD_ERROR, Message: "Unable to add suscription to provider!"}
}

func parseUserProviderToUserProviderResponse(provider *models.Provider) responses.ProviderResponse {
	return responses.ProviderResponse{
		ID:      provider.ID,
		Name:    provider.Name,
		Contact: provider.Contact,
	}
}
