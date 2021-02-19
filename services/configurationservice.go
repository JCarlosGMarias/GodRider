package services

import "godrider/infrastructures"

// ConfigurationServicer provides an access to data intended to modify
// some aspects in those applications which cosume this API.
// For instance, it provides an updated endpoint list for all methods.
type ConfigurationServicer interface {
	// GetAPIUrls should return an updated endpoint list for all API methods as k-v strings map.
	GetAPIUrls() map[string]string
}

// ConfigurationService is ConfigurationServicer's implementation struct
type ConfigurationService struct {
	apiUrlsInfrastructure infrastructures.APIUrlsInfrastructurer
}

// GetAPIUrls returns an updated endpoint list for all API methods as k-v strings map.
func (s *ConfigurationService) GetAPIUrls() map[string]string {
	apiUrls, _, _ := s.apiUrlsInfrastructure.GetAllUrls()

	result := map[string]string{}
	for _, apiURL := range apiUrls {
		result[apiURL.Key] = apiURL.Url
	}
	return result
}

// APIUrlsInfrastructure setter
func (s *ConfigurationService) APIUrlsInfrastructure(i *infrastructures.APIUrlsInfrastructurer) {
	if s.apiUrlsInfrastructure == nil {
		s.apiUrlsInfrastructure = *i
	}
}
