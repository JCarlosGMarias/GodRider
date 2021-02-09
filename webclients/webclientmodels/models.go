package webclientmodels

type ClientData struct {
	ProviderID int
	User       string
	Password   string
	Token      string
}

type Order struct {
	CustomerName     string
	Business         string
	ReceptionAddress string
	ShippingAddress  string
	ReceptionCoords  []float32
	ShippingCoords   []float32
	Amount           float32
	OrderLines       []OrderLine
}

type OrderLine struct {
	ProductName string
	Price       float32
	Quantity    uint16
}
