package services

import (
	"godrider/dtos/requests"
	"godrider/dtos/responses"
	"godrider/infrastructures"
	"godrider/infrastructures/models"
)

// UserServicer is the user's service layer
type UserServicer interface {
	// GetAllUsers should return all db user models in an UserResponse's format
	GetAllUsers() []responses.UserResponse
	// GetUserByCredentials should return an unique user based on the provided credentials (user and pass)
	GetUserByCredentials(userRq *requests.UserRequest) (responses.UserResponse, error)
	// GetUserByToken should return an unique user based on the provided credentials (token)
	GetUserByToken(userRq *requests.UserRequest) (responses.UserResponse, error)
}

// UserService is the UserServicer's implementation struct
type UserService struct {
	userInfrastructure infrastructures.UserInfrastructurer
}

// UserSrv is the UserServicer's implementation instance
var UserSrv UserServicer = &UserService{userInfrastructure: infrastructures.UsersDb}

// GetAllUsers returns all db user models in an UserResponse's format
func (service *UserService) GetAllUsers() []responses.UserResponse {
	users, count, _ := service.userInfrastructure.GetAllUsers()

	usersResponses := make([]responses.UserResponse, count)
	for _, user := range users {
		usersResponse := parseUserToUserResponse(&user)
		usersResponses = append(usersResponses, usersResponse)
	}

	return usersResponses
}

// GetUserByCredentials returns an unique user based on the provided credentials (user and pass)
func (service *UserService) GetUserByCredentials(userRq *requests.UserRequest) (responses.UserResponse, error) {
	user, err := service.userInfrastructure.GetSingleUserByUserAndPass(userRq.User, userRq.Password)
	if err != nil || user.ID == 0 {
		return responses.UserResponse{}, &responses.ErrorResponse{Code: responses.BAD_CREDENTIALS, Message: "User or password are not correct!"}
	}

	userResponse := parseUserToUserResponse(&user)
	return userResponse, nil
}

// GetUserByToken returns an unique user based on the provided credentials (token)
func (service *UserService) GetUserByToken(userRq *requests.UserRequest) (responses.UserResponse, error) {
	user, err := service.userInfrastructure.GetSingleUserByToken(userRq.Token)
	if err != nil || user.ID == 0 {
		return responses.UserResponse{}, &responses.ErrorResponse{Code: responses.BAD_TOKEN, Message: "User token not found or expired!"}
	}

	userResponse := parseUserToUserResponse(&user)
	return userResponse, nil
}

func parseUserToUserResponse(user *models.User) responses.UserResponse {
	return responses.UserResponse{
		ID:      user.ID,
		Token:   user.Token.String,
		User:    user.User,
		Name:    user.Name,
		Surname: user.Surname,
		Email:   user.Email,
		Phone:   user.Phone,
		Level:   user.Level,
	}
}
