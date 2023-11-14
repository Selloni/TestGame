package posgresql

import (
	"WB/interal"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

//type Client interface {
//	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
//	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
//	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
//	Begin(ctx context.Context) (pgx.Tx, error)
//}

// todo:config

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

func NewClient(ctx context.Context) (*pgxpool.Pool, error) {
	connConfig := "postgres://grandpat:grandpat@localhost:5432/postgres"
	//if err != nil {
	//	fmt.Errorf("Ошибка при разборе параметров подключения:%v", err)
	//}
	conn, err := pgxpool.Connect(context.Background(), connConfig)
	if err != nil {
		fmt.Errorf("Ошибка при подключении к базе данных:%v", err)
	}
	//defer conn.Close(context.Background())

	return conn, nil
}