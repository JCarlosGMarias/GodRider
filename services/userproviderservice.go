package services

import (
	"godrider/dtos/requests"
	"godrider/dtos/responses"
	"godrider/infrastructures"
	"godrider/infrastructures/models"
)

type UserProviderService struct {
	userProviderInfrastructure infrastructures.UserProviderInfrastructure
}

var UserProviderSrv = UserProviderService{
	userProviderInfrastructure: infrastructures.UserProviderDb,
}

func (service *UserProviderService) AddConnection(request *requests.UserProviderRequest) error {
	model := parseUserProviderRequestToUserProvider(request)
	model.IsActive = 1

	err := service.userProviderInfrastructure.InsertSingle(&model)
	if err == nil {
		return &responses.ErrorResponse{Code: responses.OK, Message: ""}
	}
	return &responses.ErrorResponse{Code: responses.ADD_ERROR, Message: "Unable to add suscription to provider!"}
}

func (service *UserProviderService) UpdateConnection(request *requests.UserProviderRequest) error {
	model := parseUserProviderRequestToUserProvider(request)

	err := service.userProviderInfrastructure.UpdateSingle(&model)
	if err == nil {
		return &responses.ErrorResponse{Code: responses.OK, Message: ""}
	}
	return &responses.ErrorResponse{Code: responses.ADD_ERROR, Message: "Unable to update suscription to provider!"}
}

func parseUserProviderRequestToUserProvider(request *requests.UserProviderRequest) models.UserProvider {
	model := models.UserProvider{
		UserId:     request.UserId,
		ProviderId: request.ProviderId,
	}
	if request.IsActive {
		model.IsActive = 1
	} else {
		model.IsActive = 0
	}
	return model
}
