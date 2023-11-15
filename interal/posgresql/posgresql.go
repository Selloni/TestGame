package posgresql

import (
	"WB/interal"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

func Check(ctx context.Context, conn *pgxpool.Pool, user interal.Model) (bool, error) {
	var count int
	q := fmt.Sprintf(`
			select id from %s
			where login = ($1) and password = ($2)
			`, user.Role)
	if err := conn.QueryRow(ctx, q, user.Login, user.Password).Scan(&count); err != nil {
		if err == pgx.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return count > 0, nil
}

// todo:config

func NewClient(ctx context.Context) (*pgxpool.Pool, error) {
	connConfig := "postgres://grandpat:grandpat@localhost:5432/postgres"
	conn, err := pgxpool.Connect(context.Background(), connConfig)
	if err != nil {
		fmt.Errorf("Ошибка при подключении к базе данных:%v", err)
	}

	return conn, nil
}
