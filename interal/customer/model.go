package customer

type Customer struct {
	Login    string             `json:"login,omitempty"`
	Password string             `json:"password,omitempty"`
	Money    int                `json:"money,omitempty" json:"money,omitempty"`
	Loader   []string           `json:"loader,omitempty" json:"loader,omitempty"`
	Tasks    map[string]float64 `json:"tasks,omitempty" json:"tasks,omitempty"`
}
