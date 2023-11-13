package loader

type Loader struct {
	Login    string  `json:"login,omitempty"`
	Password string  `json:"password,omitempty"`
	Weight   float64 `json:"weight,omitempty"`
	Salary   int     `json:"salary,omitempty"`
	Drunk    bool    `json:"drunk,omitempty"`
	Tired    int     `json:"tired,omitempty"`
}
