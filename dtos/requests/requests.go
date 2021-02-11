// This package groups all DTOs used for incoming API requests
package requests

type BaseRequest struct {
	Token string `json:"token"`
}

type UserRequest struct {
	User     string `json:"user"`
	Password string `json:"pass"`
	Token    string
}

type ProviderRequest struct {
	BaseRequest
	ProviderID int  `json:"provider-id"`
	IsActive   bool `json:"is-active"`
}

type UserProviderRequest struct {
	UserId     int
	ProviderId int
	IsActive   bool
}

type ApiUrlsRequest struct {
	BaseRequest
}

type OrderRequest struct {
	BaseRequest
	ProviderIDs []int `json:"provider-ids"`
}
