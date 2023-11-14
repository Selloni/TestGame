package db

import (
	"WB/interal"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
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

func CreateUser(ctx context.Context, conn *pgx.Conn, user interal.Model) error {
	q := `
		insert into customer
			(login, password, money)
		values 
		    ($1,$2,$3)
-- 		returning id
		`

	err := conn.QueryRow(ctx, q, user.Login, user.Password, user.Customer.Money)
	if err != nil {
		return fmt.Errorf("не удалось создать заказчика")
	}
	fmt.Println(q)
	return nil
}

func Check(ctx context.Context, conn *pgx.Conn, user interal.Model) (bool, error) {
	var count int
	q := `
			select id from customer
			where login = ($1)
			`
	if err := conn.QueryRow(ctx, q, user.Login).Scan(&count); err != nil {
		if err == pgx.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return count > 0, nil
}

//func GetTask(ctx context.Context) (map[string]float64, error) {
//
//}
//
//func GetInfo(ctx context.Context) (*interal.Model, error) {
//
//}
