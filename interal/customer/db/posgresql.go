package db

import (
	"WB/interal"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
)

func CreateUser(ctx context.Context, user interal.Model, conn *pgx.Conn) error {
	q := `
		insert into customer
			(login, password, money)
		values 
		    ($1,$2,$3)
		returning id
		`
	err := conn.QueryRow(ctx, q, user.Login, user.Password, user.Customer.Money)
	if err != nil {
		return fmt.Errorf("не удалось создать заказчика")
	}
	return nil
}

func GetTask(ctx context.Context) (map[string]float64, error) {

}

func GetInfo(ctx context.Context) (*interal.Model, error) {

}
