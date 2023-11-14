package db

import (
	"WB/interal"
	"WB/interal/customer"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

//type repository struct {
//	sql *pgx.Conn
//}
//
//func NewRepository(conn *pgx.Conn) *repository {
//	return &repository{
//		sql: conn,
//	}
//}

func CreateUser(ctx context.Context, conn *pgxpool.Pool, user interal.Model) error {
	q := `
		insert into customer
			(login, password, money)
		values 
		    ($1,$2,$3)
		`
	err := conn.QueryRow(ctx, q, user.Login, user.Password, user.Customer.Money).Scan()
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil
		}
		return fmt.Errorf("не удалось создать заказчика")
	}
	return nil
}

//func GetTask(ctx context.Context) (map[string]float64, error) {
//
//}
//

func GetInfo(ctx context.Context, conn *pgxpool.Pool, login string) (*customer.Customer, error) {
	q := `
		select money from customer where login = $1
	`
	var cust customer.Customer
	err := conn.QueryRow(ctx, q, login).Scan(&cust.Money)
	if err != nil {
		return &customer.Customer{}, err
	}
	return &cust, nil
}
