package services

import (
	"godrider/dtos/requests"
	"godrider/dtos/responses"
	"godrider/infrastructures"
	"godrider/infrastructures/models"
)

// UserServicer is the user's service layer
type UserServicer interface {
	// GetAll should return all db user models in an UserResponse's format
	GetAll() []responses.UserResponse
	// GetByCredentials should return an unique user based on the provided credentials (user and pass)
	GetByCredentials(userRq *requests.UserRequest) (responses.UserResponse, error)
	// GetByToken should return an unique user based on the provided credentials (token)
	GetByToken(userRq *requests.UserRequest) (responses.UserResponse, error)
}

// UserService is the UserServicer's implementation struct
type UserService struct {
	userIstruct infrastructures.UserInfrastructurer
}

// GetAll returns all db user models in an UserResponse's format
func (s *UserService) GetAll() []responses.UserResponse {
	users, count, _ := s.userIstruct.GetAll()

	usersRs := make([]responses.UserResponse, count)
	for i, user := range users {
		usersRs[i] = parseUserToUserResponse(&user)
	}

	return usersRs
}

// GetByCredentials returns an unique user based on the provided credentials (user and pass)
func (s *UserService) GetByCredentials(userRq *requests.UserRequest) (responses.UserResponse, error) {
	user, err := s.userIstruct.GetSingleByUserAndPass(userRq.User, userRq.Password)
	if err != nil || user.ID == 0 {
		return responses.UserResponse{}, &responses.ErrorResponse{Code: responses.BAD_CREDENTIALS, Message: "User or password are not correct!"}
	}

	userRs := parseUserToUserResponse(&user)
	return userRs, nil
}

// GetByToken returns an unique user based on the provided credentials (token)
func (s *UserService) GetByToken(userRq *requests.UserRequest) (responses.UserResponse, error) {
	user, err := s.userIstruct.GetSingleByToken(userRq.Token)
	if err != nil || user.ID == 0 {
		return responses.UserResponse{}, &responses.ErrorResponse{Code: responses.BAD_TOKEN, Message: "User token not found or expired!"}
	}

	userRs := parseUserToUserResponse(&user)
	return userRs, nil
}

// UserInfrastructure setter
func (s *UserService) UserInfrastructure(i *infrastructures.UserInfrastructurer) {
	if s.userIstruct == nil {
		s.userIstruct = *i
	}
}

func parseUserToUserResponse(u *models.User) responses.UserResponse {
	return responses.UserResponse{
		ID:      u.ID,
		Token:   u.Token.String,
		User:    u.User,
		Name:    u.Name,
		Surname: u.Surname,
		Email:   u.Email,
		Phone:   u.Phone,
		Level:   u.Level,
	}
}
