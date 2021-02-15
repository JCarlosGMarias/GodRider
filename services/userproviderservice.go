package services

import (
	"godrider/dtos/requests"
	"godrider/dtos/responses"
	"godrider/infrastructures"
	"godrider/infrastructures/models"
)

// UserProviderServicer is the userprovider's service layer
type UserProviderServicer interface {
	// AddConnection should create a new userprovider connection or return a custom service error
	AddConnection(request *requests.UserProviderRequest) error
	// UpdateConnection should change an existent connection's active field successfully or return a custom error
	UpdateConnection(request *requests.UserProviderRequest) error
}

// UserProviderService is UserProviderServicer's implementation struct
type UserProviderService struct {
	userProviderInfrastructure infrastructures.UserProviderInfrastructurer
}

// AddConnection creates a new userprovider connection or returns a custom service error
func (service *UserProviderService) AddConnection(request *requests.UserProviderRequest) error {
	model := parseUserProviderRequestToUserProvider(request)
	model.IsActive = 1

	err := service.userProviderInfrastructure.InsertSingle(&model)
	if err == nil {
		return &responses.ErrorResponse{Code: responses.OK, Message: ""}
	}
	return &responses.ErrorResponse{Code: responses.ADD_ERROR, Message: "Unable to add suscription to provider!"}
}

// UpdateConnection changes an existent connection's active field successfully or returns a custom error
func (service *UserProviderService) UpdateConnection(request *requests.UserProviderRequest) error {
	model := parseUserProviderRequestToUserProvider(request)

	err := service.userProviderInfrastructure.UpdateSingle(&model)
	if err == nil {
		return &responses.ErrorResponse{Code: responses.OK, Message: ""}
	}
	return &responses.ErrorResponse{Code: responses.ADD_ERROR, Message: "Unable to update suscription to provider!"}
}

// UserProviderInfrastructure setter
func (service *UserProviderService) UserProviderInfrastructure(istruct *infrastructures.UserProviderInfrastructurer) {
	if service.userProviderInfrastructure == nil {
		service.userProviderInfrastructure = *istruct
	}
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
