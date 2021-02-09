package controllers

import (
	"encoding/json"
	"net/http"

	"godrider/dtos/requests"
	"godrider/helpers"
	"godrider/services"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {
	if helpers.IsValidMethod(w, r, []string{http.MethodPost}) {
		var orderRq requests.OrderRequest
		helpers.ParseBody(r.Body, &orderRq)

		if helpers.IsValidToken(w, r, orderRq.Token) {
			json.NewEncoder(w).Encode(services.OrderSrv.GetOrders(&orderRq))
		}
	}
}
