package interal

import (
	"WB/interal/customer"
	"WB/interal/loader"
)

type Model struct {
	Login    string             `json:"login,omitempty"`
	Password string             `json:"password,omitempty"`
	Role     string             `json:"role,omitempty"`
	Customer *customer.Customer `json:"-"`
	Loader   *loader.Loader     `json:"-"`
}

type StartGameRequest struct {
	TaskId  int   `json:"taskId"`
	Loaders []int `json:"loaders"`
}
