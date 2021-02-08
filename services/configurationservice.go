package services

import "godrider/infrastructures"

// ConfigurationService provides an access to data intended to modify
// some aspects in those applications which cosume this API.
// For instance, it provides an updated endpoint list for all methods.
type ConfigurationService struct {
	apiUrlsInfrastructure infrastructures.ApiUrlsInfrastructure
}

var ConfigSrv = ConfigurationService{
	apiUrlsInfrastructure: infrastructures.ApiUrlsDb,
}

// GetApiUrls returns an updated endpoint list for all API methods
// as k-v strings map.
func (service *ConfigurationService) GetApiUrls() map[string]string {
	apiUrls, _, _ := service.apiUrlsInfrastructure.GetAllUrls()

	result := map[string]string{}
	for _, apiUrl := range apiUrls {
		result[apiUrl.Key] = apiUrl.Url
	}
	return result
}
