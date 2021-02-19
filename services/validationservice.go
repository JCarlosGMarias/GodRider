package services

import (
	"fmt"
	"godrider/dtos/requests"
	"godrider/dtos/responses"
)

// ValidationServicer defines a group of common validation tools across the API
type ValidationServicer interface {
	// ValidateMethod should read the Request's Method field, search it into allowedMethods and determine if is an allowed method for the calling controller endpoint
	ValidateMethod(method string, allowedMethods []string) error
	// ValidateToken should determine if the request token belongs to an existing or active user
	ValidateToken(token string) error
}

// ValidationService is ValidationServicer's implemenation struct
type ValidationService struct {
	userSrv UserServicer
}

// ValidateMethod reads the Request's Method field, searches it into allowedMethods and determines if is an allowed method for the calling controller endpoint
func (s *ValidationService) ValidateMethod(method string, allowedMethods []string) error {
	for _, m := range allowedMethods {
		if method == m {
			return nil
		}
	}

	message := fmt.Sprintf("Incoming request with method %s not allowed!", method)
	return &responses.ErrorResponse{Code: responses.METHOD_UNAUTHORIZED, Message: message}
}

// ValidateToken determines if the request token belongs to an existing or active user
func (s *ValidationService) ValidateToken(token string) error {
	_, err := s.userSrv.GetByToken(&requests.UserRequest{Token: token})

	if err == nil {
		return nil
	}
	return &responses.ErrorResponse{Code: responses.BAD_TOKEN, Message: "User token not found or expired!"}
}

// UserSrv setter
func (s *ValidationService) UserSrv(srv *UserServicer) {
	if s.userSrv == nil {
		s.userSrv = *srv
	}
}
