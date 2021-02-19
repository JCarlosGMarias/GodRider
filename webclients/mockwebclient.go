package webclients

import (
	"fmt"
	"godrider/webclients/webclientmodels"
	"math/rand"
)

type mockWebClient struct {
	mockUser     string
	mockPassword string
}

func (client *mockWebClient) Login(data *webclientmodels.ClientData) (bool, error) {
	if data.User == client.mockUser && data.Password == client.mockPassword {
		data.Token = "jhknlfsdv786y3421rhiblyu"
		return true, nil
	}
	return false, fmt.Errorf("Login error: user or password incorrect")
}

func (client *mockWebClient) GetOrders(c chan []webclientmodels.Order) error {
	// Here we should make our request to the external webservice and perform
	// all needed tasks in order to get a proper data response for our calling services...
	randomPrice := 10 + rand.Float32()*(100-10)
	c <- []webclientmodels.Order{
		{
			CustomerName:     "Andrew",
			Business:         "PizzaPlanet",
			ReceptionAddress: "C/Borricondeabajo, 73, 1A",
			ShippingAddress:  "C/Borricondearriba, 10, 1D",
			ReceptionCoords:  []float32{-27.967917, 153.419083},
			ShippingCoords:   []float32{30.541639, 47.825444},
			// Amount:           44.85,
			Amount: randomPrice,
			OrderLines: []webclientmodels.OrderLine{
				{
					ProductName: "Prosciutto parmigiano pineapple pizza",
					Price:       5.95,
					Quantity:    3,
				},
			},
		},
		{
			CustomerName:     "Concha",
			Business:         "Casa Manolo",
			ReceptionAddress: "C/Borricondearriba, 10, 1D",
			ShippingAddress:  "C/Desengaño, 21, 1 Izda",
			ReceptionCoords:  []float32{50.844028, -0.172361},
			ShippingCoords:   []float32{37.660611, -116.028361},
			Amount:           44.85,
			OrderLines: []webclientmodels.OrderLine{
				{
					ProductName: "Callos a la riojana",
					Price:       14.95,
					Quantity:    10,
				},
			},
		},
	}
	return nil
}
