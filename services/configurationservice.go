package services

// ConfigurationService provides an access to data intended to modify
// some aspects in those applications which cosume this API.
// For instance, it provides an updated endpoint list for all methods.
type ConfigurationService struct {
	urlsDb map[string]string
}

var ConfigSrv = ConfigurationService{
	urlsDb: map[string]string{
		"LoginUrl":             "/api/Login",
		"GetApiUrlsUrl":        "/api/GetApiUrls",
		"GetProvidersUrl":      "/api/GetProviders",
		"ConnectToProviderUrl": "/api/ConnectToProvider",
		"RevokeProviderUrl":    "/api/RevokeProvider",
		"GetOrdersUrl":         "/api/GetOrders",
		"AssignOrderUrl":       "/api/AssignOrder",
	},
}

// GetApiUrls returns an updated endpoint list for all API methods
// as k-v strings map.
func (service *ConfigurationService) GetApiUrls() map[string]string {
	return service.urlsDb
}
