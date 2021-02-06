package requests

type BaseRequest struct {
	Token string `json:"token"`
}

type UserRequest struct {
	User     string `json:"user"`
	Password string `json:"pass"`
}

type ProviderRequest struct {
	ProviderID int `json:"provider-id"`
	BaseRequest
}

type ApiUrlsRequest struct {
	BaseRequest
}
