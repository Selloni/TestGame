package loader

import (
	"math/rand"
	"time"
)

type Loader struct {
	Weight float64 `json:"weight,omitempty"`
	Salary int     `json:"salary,omitempty"`
	Drunk  bool    `json:"drunk,omitempty"`
	Tired  int     `json:"tired,omitempty"`
}

func GenerateLoader() *Loader {
	drunk := func() bool {
		rand.Seed(time.Now().UnixNano())
		return rand.Intn(2) == 1
	}()
	loader := Loader{
		Weight: 5 + rand.Float64()*(25),
		Salary: rand.Intn(20000) + 10000,
		Drunk:  drunk,
		Tired:  0,
	}
	return &loader
}
