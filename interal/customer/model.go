package customer

import "WB/interal/loader"

type Customer struct {
	Money  int                `json:"money,omitempty" json:"money,omitempty"`
	Loader []loader.Loader    `json:"loader,omitempty" json:"loader,omitempty"`
	Tasks  map[string]float64 `json:"tasks,omitempty" json:"tasks,omitempty"`
}
