package controllers

import (
	"net/http"

	"godrider/dtos/requests"
	"godrider/helpers"
	"godrider/services"
)

type ConfigurationControllerer interface {
	GetRoutes() map[string]string
	GetApiUrls(w http.ResponseWriter, r *http.Request)
}

type ConfigurationController struct {
	validationSrv services.ValidationServicer
	configSrv     services.ConfigurationServicer
}

func (c *ConfigurationController) GetRoutes() map[string]string {
	return c.configSrv.GetAPIUrls()
}

func (c *ConfigurationController) GetApiUrls(w http.ResponseWriter, r *http.Request) {
	if err := c.validationSrv.ValidateMethod(r.Method, []string{http.MethodPost}); err == nil {
		var apiUrlsRq requests.ApiUrlsRequest
		helpers.ParseBody(r.Body, &apiUrlsRq)

		e := helpers.SetCommonResponseEncoder(&w)
		if err := c.validationSrv.ValidateToken(apiUrlsRq.Token); err == nil {
			e.Encode(c.configSrv.GetAPIUrls())
		} else {
			e.Encode(err)
		}
	}
}

// ValidationSrv setter
func (c *ConfigurationController) ValidationSrv(s *services.ValidationServicer) {
	if c.validationSrv == nil {
		c.validationSrv = *s
	}
}

// ConfigSrv setter
func (c *ConfigurationController) ConfigSrv(s *services.ConfigurationServicer) {
	if c.configSrv == nil {
		c.configSrv = *s
	}
}
