package services

import (
	"godrider/dtos/requests"
	"godrider/dtos/responses"
	"godrider/infrastructures"
	"godrider/webclients"
	"godrider/webclients/webclientmodels"
)

type OrderService struct {
	providerInfrastructure infrastructures.ProvidersInfrastructure
	factory                webclients.WebClientFactory
}

var OrderSrv = OrderService{
	providerInfrastructure: infrastructures.ProvidersDb,
	factory:                webclients.ClientFactory,
}

func (service *OrderService) GetOrders(request *requests.OrderRequest) []responses.OrderResponse {
	orderResults := make([]responses.OrderResponse, 0)

	providers, err := service.providerInfrastructure.GetManyProvidersByIds(request.ProviderIDs)
	if err != nil {
		err = &responses.ErrorResponse{Code: responses.REGISTER_NOT_FOUND, Message: "Unable to recover orders from client!"}
		return make([]responses.OrderResponse, 0)
	}

	for _, provider := range providers {
		clientData := &webclientmodels.ClientData{
			ProviderID: provider.ID,
			User:       provider.Name,
			Password:   provider.Contact,
			Token:      provider.Contact,
		}

		webClient, err := service.factory.GetClient(clientData)
		if err != nil {
			err = &responses.ErrorResponse{Code: responses.WEBSERVICE_CONNECTION_FAILURE, Message: "Unable to recover orders from client!"}
			return make([]responses.OrderResponse, 0)
		}

		orders, _ := webClient.GetOrders()

		for _, order := range orders {
			orderResults = append(orderResults, service.parseOrderToOrderResponse(&order))
		}
	}

	return orderResults
}

func (service *OrderService) parseOrderToOrderResponse(order *webclientmodels.Order) responses.OrderResponse {
	response := responses.OrderResponse{
		CustomerName:     order.CustomerName,
		Business:         order.Business,
		ReceptionAddress: order.ReceptionAddress,
		ShippingAddress:  order.ShippingAddress,
		ReceptionCoords:  order.ReceptionCoords,
		ShippingCoords:   order.ShippingCoords,
		Amount:           order.Amount,
		OrderLines:       make([]responses.OrderLineResponse, 0),
	}

	for _, orderLine := range order.OrderLines {
		response.OrderLines = append(response.OrderLines, service.parseOrderLineToOrderLineResponse(&orderLine))
	}
	return response
}

func (service *OrderService) parseOrderLineToOrderLineResponse(orderLine *webclientmodels.OrderLine) responses.OrderLineResponse {
	return responses.OrderLineResponse{
		ProductName: orderLine.ProductName,
		Price:       orderLine.Price,
		Quantity:    orderLine.Quantity,
	}
}
