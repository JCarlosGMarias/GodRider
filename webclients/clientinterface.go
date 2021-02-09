package webclients

import "godrider/webclients/webclientmodels"

type WebClient interface {
	// Login defines a method in which we perform a login
	// attempt to the pertinent external webservice.
	//
	// Also, we could store some temporal info such as tokens
	// or similar, it depends on the webservice logic flow.
	// In general, use bool for a simple success flag and error
	// for details about a login failure.
	Login(data *webclientmodels.ClientData) (bool, error)
	// GetOrders defines a method in which we fetch all
	// available orders from the given webservice and make
	// some parsing operations in order to return an order
	// slice.
	//
	// If the webservice needs some kind of data in order to filter
	// results, we should provide the most wide option available.
	// However, a new method signature with filter options is coming,
	// so this method should be used only when we want the most wide
	// response.
	GetOrders() ([]webclientmodels.Order, error)
}
