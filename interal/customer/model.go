package customer

import (
	"math/rand"
)

type Customer struct {
	Money int
}

func GenerateCustomer() *Customer {
	customer := Customer{
		Money: rand.Intn(90000+1) + 10000,
	}
	return &customer
}
