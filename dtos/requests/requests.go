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
	ProviderID int `json:"provider-id"`
	BaseRequest
}

type UserProviderRequest struct {
	UserId     int
	ProviderId int
	IsActive   bool
}

type ApiUrlsRequest struct {
	BaseRequest
}
