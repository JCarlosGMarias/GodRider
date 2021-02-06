package services

import (
	"godrider/dtos/requests"
	"godrider/dtos/responses"
	"godrider/infrastructures"
	"godrider/infrastructures/models"
)

type UserService struct {
	userInfrastructure infrastructures.UsersInfrastructure
}

var UserSrv = UserService{
	userInfrastructure: infrastructures.UsersDb,
}

func (service *UserService) GetAllUsers() []responses.UserResponse {
	users, count, _ := service.userInfrastructure.GetAllUsers()

	usersResponses := make([]responses.UserResponse, count)
	for _, user := range users {
		usersResponse := parseUserToUserResponse(&user)
		usersResponses = append(usersResponses, usersResponse)
	}

	return usersResponses
}

func (service *UserService) GetUserByCredentials(userRq *requests.UserRequest) (responses.UserResponse, responses.ErrorResponse) {
	user, err := service.userInfrastructure.GetSingleUserByUserAndPass(userRq.User, userRq.Password)
	if err != nil || user.ID == 0 {
		return responses.UserResponse{}, responses.ErrorResponse{Code: 2, Message: "User or password are not correct!"}
	}

	userResponse := parseUserToUserResponse(&user)
	return userResponse, responses.ErrorResponse{}
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
