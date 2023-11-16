package customer

import (
	"WB/interal/loader"
	"math/rand"
)

type Customer struct {
	Money  int             `json:"money,omitempty" json:"money,omitempty"`
	Loader []loader.Loader `json:"loader,omitempty" json:"loader,omitempty"`
	Tasks  map[string]int  `json:"tasks,omitempty" json:"tasks,omitempty"`
}

func GenerateCustomer() *Customer {
	// сгенерировать рандомное число
	customer := Customer{
		Money:  rand.Intn(90000+1) + 10000,
		Loader: nil,
		Tasks:  nil,
	}
	return &customer
}
