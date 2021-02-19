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
func (s *UserProviderService) AddConnection(request *requests.UserProviderRequest) error {
	model := parseUserProviderRequestToUserProvider(request)
	model.IsActive = 1

	err := s.userProviderInfrastructure.InsertSingle(&model)
	if err == nil {
		return &responses.ErrorResponse{Code: responses.OK, Message: ""}
	}
	return &responses.ErrorResponse{Code: responses.ADD_ERROR, Message: "Unable to add suscription to provider!"}
}

// UpdateConnection changes an existent connection's active field successfully or returns a custom error
func (s *UserProviderService) UpdateConnection(request *requests.UserProviderRequest) error {
	model := parseUserProviderRequestToUserProvider(request)

	err := s.userProviderInfrastructure.UpdateSingle(&model)
	if err == nil {
		return &responses.ErrorResponse{Code: responses.OK, Message: ""}
	}
	return &responses.ErrorResponse{Code: responses.ADD_ERROR, Message: "Unable to update suscription to provider!"}
}

// UserProviderInfrastructure setter
func (s *UserProviderService) UserProviderInfrastructure(i *infrastructures.UserProviderInfrastructurer) {
	if s.userProviderInfrastructure == nil {
		s.userProviderInfrastructure = *i
	}
}

func parseUserProviderRequestToUserProvider(r *requests.UserProviderRequest) models.UserProvider {
	model := models.UserProvider{
		UserId:     r.UserId,
		ProviderId: r.ProviderId,
	}
	if r.IsActive {
		model.IsActive = 1
	} else {
		model.IsActive = 0
	}
	return model
}
