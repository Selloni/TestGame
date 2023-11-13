package customer

import (
	"context"
)

type Repository interface {
	CreateUser(ctx context.Context, customer *Customer) error
	GetTask(ctx context.Context) (map[string]float64, error)
	GetInfo(ctx context.Context) (*Customer, error)
}
