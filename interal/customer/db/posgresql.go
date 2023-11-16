package db

import (
	"WB/interal"
	"WB/interal/customer"
	"WB/interal/task"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

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

func GetInfo(ctx context.Context, conn *pgxpool.Pool, login string) (*customer.Customer, error) {
	q := `
		select money from customer where login = $1
	`
	var cust customer.Customer
	err := conn.QueryRow(ctx, q, login).Scan(&cust.Money)
	if err != nil {
		return nil, err
	}
	return &cust, nil
}

func GetAllTask(ctx context.Context, conn *pgxpool.Pool) ([]task.Task, error) {
	arrTask := make([]task.Task, 0)
	q := `
		select id, name, weight from task where done = false
		`
	rows, err := conn.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t task.Task

		if err := rows.Scan(&t.Id, &t.Name, &t.Weight); err != nil {
			return nil, err
		}
		arrTask = append(arrTask, t)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return arrTask, nil
}

func UpdateCustomer(ctx context.Context, conn *pgxpool.Pool, user interal.Model) error {
	q := `
		update customer
			set money = $1
			where login = $2
	`
	_, err := conn.Exec(ctx, q, user.Customer.Money, user.Login)
	if err != nil {
		return err
	}
	return nil
}
