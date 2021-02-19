package services

import (
	"fmt"
	"godrider/dtos/requests"
	"godrider/dtos/responses"
	"godrider/infrastructures"
	"godrider/webclients"
	"godrider/webclients/webclientmodels"
	"log"
)

// OrderServicer provides access to get customer's orders and edit its status
type OrderServicer interface {
	// GetOrders should identify the requested providers and return a slice with all available customer's orders
	GetOrders(request *requests.OrderRequest) []responses.OrderResponse
}

// OrderService is OrderServicer's implementation struct
type OrderService struct {
	providerInfrastructure infrastructures.ProviderInfrastructurer
	factory                webclients.WebClientFactorier
}

// GetOrders identifies the requested providers and returns a slice with all available customer's orders
func (s *OrderService) GetOrders(request *requests.OrderRequest) []responses.OrderResponse {
	providers, err := s.providerInfrastructure.GetManyByIds(request.ProviderIDs)
	if err != nil {
		err = &responses.ErrorResponse{Code: responses.REGISTER_NOT_FOUND, Message: "Unable to recover orders from client!"}
		return []responses.OrderResponse{}
	}

	orderResults := make([]responses.OrderResponse, 0)
	c := make(chan []webclientmodels.Order, len(providers))
	for _, provider := range providers {
		clientData := &webclientmodels.ClientData{
			ProviderID: provider.ID,
			User:       provider.Name,
			Password:   provider.Contact,
			Token:      provider.Contact,
		}

		webClient, err := s.factory.GetClient(clientData)
		if err != nil {
			message := fmt.Sprintf("Unable to recover orders from client with ID %d!", clientData.ProviderID)
			err = &responses.ErrorResponse{Code: responses.WEBSERVICE_CONNECTION_FAILURE, Message: message}
			log.Print(err)
			continue
		}

		go webClient.GetOrders(c)
	}

	for i := 0; i < len(providers); i++ {
		orders, ok := <-c
		if !ok {
			break
		}
		for _, order := range orders {
			orderResults = append(orderResults, parseOrderToOrderResponse(&order))
		}
	}
	close(c)

	return orderResults
}

// ProviderInfrastructure setter
func (s *OrderService) ProviderInfrastructure(i *infrastructures.ProviderInfrastructurer) {
	if s.providerInfrastructure == nil {
		s.providerInfrastructure = *i
	}
}

// Factory setter
func (s *OrderService) Factory(f *webclients.WebClientFactorier) {
	if s.factory == nil {
		s.factory = *f
	}
}

func parseOrderToOrderResponse(o *webclientmodels.Order) responses.OrderResponse {
	response := responses.OrderResponse{
		CustomerName:     o.CustomerName,
		Business:         o.Business,
		ReceptionAddress: o.ReceptionAddress,
		ShippingAddress:  o.ShippingAddress,
		ReceptionCoords:  o.ReceptionCoords,
		ShippingCoords:   o.ShippingCoords,
		Amount:           o.Amount,
		OrderLines:       make([]responses.OrderLineResponse, 0),
	}

	for _, l := range o.OrderLines {
		response.OrderLines = append(response.OrderLines, parseOrderLineToOrderLineResponse(&l))
	}
	return response
}

func parseOrderLineToOrderLineResponse(l *webclientmodels.OrderLine) responses.OrderLineResponse {
	return responses.OrderLineResponse{
		ProductName: l.ProductName,
		Price:       l.Price,
		Quantity:    l.Quantity,
	}
}
