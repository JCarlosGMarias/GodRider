package controllers

import (
	"net/http"

	"godrider/dtos/requests"
	"godrider/helpers"
	"godrider/services"
)

// ProviderControllerer
type ProviderControllerer interface {
	// GetProviders should return a complete list of providers
	GetProviders(w http.ResponseWriter, r *http.Request)
	ConnectToProvider(w http.ResponseWriter, r *http.Request)
}

// ProviderController
type ProviderController struct {
	validationSrv   services.ValidationServicer
	userSrv         services.UserServicer
	providerSrv     services.ProviderServicer
	userProviderSrv services.UserProviderServicer
}

// GetProviders returns a complete list of providers
func (c *ProviderController) GetProviders(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		helpers.SetOptionsResponseEncoder(&w)
		return
	}

	if err := c.validationSrv.ValidateMethod(r.Method, []string{http.MethodPost}); err == nil {
		var providerRq requests.ProviderRequest
		helpers.ParseBody(r.Body, &providerRq)

		if err := c.validationSrv.ValidateToken(providerRq.Token); err == nil {
			e := helpers.SetCommonResponseEncoder(&w)
			e.Encode(c.providerSrv.GetAll())
		}
	}
}

func (c *ProviderController) ConnectToProvider(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		helpers.SetOptionsResponseEncoder(&w)
		return
	}

	if err := c.validationSrv.ValidateMethod(r.Method, []string{http.MethodPost, http.MethodPut}); err == nil {
		var providerRq requests.ProviderRequest
		helpers.ParseBody(r.Body, &providerRq)

		if err := c.validationSrv.ValidateToken(providerRq.Token); err == nil {
			user, _ := c.userSrv.GetByToken(&requests.UserRequest{Token: providerRq.Token})

			request := requests.UserProviderRequest{
				UserId:     user.ID,
				ProviderId: providerRq.ProviderID,
				IsActive:   providerRq.IsActive,
			}

			var err error
			if r.Method == http.MethodPost {
				err = c.userProviderSrv.AddConnection(&request)
			} else if r.Method == http.MethodPut {
				err = c.userProviderSrv.UpdateConnection(&request)
			}
			e := helpers.SetCommonResponseEncoder(&w)
			e.Encode(err)
		}
	}
}

// ValidationSrv setter
func (c *ProviderController) ValidationSrv(s *services.ValidationServicer) {
	if c.validationSrv == nil {
		c.validationSrv = *s
	}
}

// UserSrv setter
func (c *ProviderController) UserSrv(srv *services.UserServicer) {
	if c.userSrv == nil {
		c.userSrv = *srv
	}
}

// ProviderSrv setter
func (c *ProviderController) ProviderSrv(srv *services.ProviderServicer) {
	if c.providerSrv == nil {
		c.providerSrv = *srv
	}
}

// UserProviderSrv setter
func (c *ProviderController) UserProviderSrv(srv *services.UserProviderServicer) {
	if c.userProviderSrv == nil {
		c.userProviderSrv = *srv
	}
}
