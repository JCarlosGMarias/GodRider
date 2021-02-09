package responses

type ErrorResponse struct {
	Code    ErrorCode
	Message string
}

type UserResponse struct {
	ID                int                `json:"-"`
	Token             string             `json:"token"`
	User              string             `json:"user"`
	Password          string             `json:"-"`
	Name              string             `json:"name"`
	Surname           string             `json:"surname"`
	Email             string             `json:"email"`
	Phone             string             `json:"phone"`
	AssignedProviders []ProviderResponse `json:"assigned-providers"`
	Level             int                `json:"-"`
}

type ProviderResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Contact string `json:"contact"`
}

type OrderResponse struct {
	CustomerName     string
	Business         string
	ReceptionAddress string
	ShippingAddress  string
	ReceptionCoords  []float32
	ShippingCoords   []float32
	Amount           float32
	OrderLines       []OrderLineResponse
}

type OrderLineResponse struct {
	ProductName string
	Price       float32
	Quantity    uint16
}
