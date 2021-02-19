package controllers

import (
	"net/http"

	"godrider/dtos/requests"
	"godrider/helpers"
	"godrider/services"
)

type OrderControllerer interface {
	GetOrders(w http.ResponseWriter, r *http.Request)
}

type OrderController struct {
	validationSrv services.ValidationServicer
	orderSrv      services.OrderServicer
}

func (c *OrderController) GetOrders(w http.ResponseWriter, r *http.Request) {
	if err := c.validationSrv.ValidateMethod(r.Method, []string{http.MethodPost}); err == nil {
		var orderRq requests.OrderRequest
		helpers.ParseBody(r.Body, &orderRq)

		if err := c.validationSrv.ValidateToken(orderRq.Token); err == nil {
			e := helpers.SetCommonResponseEncoder(&w)
			e.Encode(c.orderSrv.GetOrders(&orderRq))
		}
	}
}

// ValidationSrv setter
func (c *OrderController) ValidationSrv(s *services.ValidationServicer) {
	if c.validationSrv == nil {
		c.validationSrv = *s
	}
}

// OrderSrv setter
func (c *OrderController) OrderSrv(s *services.OrderServicer) {
	if c.orderSrv == nil {
		c.orderSrv = *s
	}
}
