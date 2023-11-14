package db

import (
	"WB/interal"
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

//func Check(ctx context.Context, conn *pgxpool.Pool, user interal.Model) (bool, error) {
//	var count int
//	q := `
//			select id from loader
//			where login = ($1)
//			`
//	if err := conn.QueryRow(ctx, q, user.Login).Scan(&count); err != nil {
//		if err == pgx.ErrNoRows {
//			return false, nil
//		}
//		return false, err
//	}
//	return count > 0, nil
//}
