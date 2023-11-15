package task

type Task struct {
	Id     int    `json:"-"`
	Name   string `json:"name,omitempty"`
	Weight int    `json:"weight,omitempty"`
	Done   bool   `json:"done,omitempty"`
}
