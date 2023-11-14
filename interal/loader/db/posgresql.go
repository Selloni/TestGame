package db

import (
	"WB/interal"
	"WB/interal/loader"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

func CreateUser(ctx context.Context, conn *pgxpool.Pool, user interal.Model) error {
	q := `
		insert into loader
			(login, password, weight, money, drunk)
		values 
		    ($1,$2,$3,$4,$5)
-- 		returning id
		`
	err := conn.QueryRow(ctx, q, user.Login, user.Password, user.Loader.Weight, user.Loader.Salary, user.Loader.Drunk).Scan()
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil
		}
		return fmt.Errorf("не удалось создать грузчика")
	}
	return nil
}

func GetInfo(ctx context.Context, conn *pgxpool.Pool, login string) ([]loader.Loader, error) {
	q := `
		select weight, money, drunk, tired from loader where login = $1
	`

	rows, err := conn.Query(ctx, q, login)

	loaderInfo := make([]loader.Loader, 0)

	for rows.Next() {
		var load loader.Loader
		if err = rows.Scan(&load.Weight, &load.Salary, &load.Drunk, &load.Tired); err != nil {
			return nil, err
		}
		loaderInfo = append(loaderInfo, load)
	}
	return loaderInfo, nil
}
