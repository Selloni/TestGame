package loader

import (
	"math/rand"
	"time"
)

type Loader struct {
	Id     int  `json:"-"`
	Weight int  `json:"weight,omitempty"`
	Salary int  `json:"salary,omitempty"`
	Drunk  bool `json:"drunk,omitempty"`
	Tired  int  `json:"tired,omitempty"`
}

func GenerateLoader() *Loader {
	drunk := func() bool {
		rand.Seed(time.Now().UnixNano())
		return rand.Intn(2) == 1
	}()
	loader := Loader{
		Weight: rand.Intn(25) + 5,
		Salary: rand.Intn(20000) + 10000,
		Drunk:  drunk,
		Tired:  0,
	}
	return &loader
}
