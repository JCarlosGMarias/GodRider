package controllers

import (
	"net/http"

	"godrider/dtos/requests"
	"godrider/helpers"
	"godrider/services"
)

type UserControllerer interface {
	Login(w http.ResponseWriter, r *http.Request)
}

type UserController struct {
	validationSrv services.ValidationServicer
	userSrv       services.UserServicer
}

func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	if err := c.validationSrv.ValidateMethod(r.Method, []string{http.MethodPost}); err == nil {
		var userReq requests.UserRequest
		helpers.ParseBody(r.Body, &userReq)

		user, errorRs := c.userSrv.GetByCredentials(&userReq)
		e := helpers.SetCommonResponseEncoder(&w)
		if errorRs == nil {
			e.Encode(user)
		} else {
			e.Encode(errorRs)
		}
	}
}

// ValidationSrv setter
func (c *UserController) ValidationSrv(s *services.ValidationServicer) {
	if c.validationSrv == nil {
		c.validationSrv = *s
	}
}

// UserSrv setter
func (c *UserController) UserSrv(s *services.UserServicer) {
	if c.userSrv == nil {
		c.userSrv = *s
	}
}
