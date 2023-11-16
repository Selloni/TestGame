package interal

import (
	"WB/interal/customer"
	"WB/interal/loader"
)

type Model struct {
	Login    string             `json:"login"`
	Password string             `json:"password"`
	Role     string             `json:"role"`
	Customer *customer.Customer `json:"-"`
	Loader   *loader.Loader     `json:"-"`
}

type StartGameRequest struct {
	TaskId  int   `json:"taskId"`
	Loaders []int `json:"loaders"`
}
