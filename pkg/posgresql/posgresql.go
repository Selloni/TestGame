package posgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
)

//type Client interface {
//	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
//	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
//	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
//	Begin(ctx context.Context) (pgx.Tx, error)
//}

// todo:config

func NewClient(ctx context.Context) (*pgx.Conn, error) {
	connConfig, err := pgx.ParseConfig("postgres://grandpat:grandpat@localhost:5432/postgres")
	if err != nil {
		fmt.Errorf("Ошибка при разборе параметров подключения:%v", err)
	}
	conn, err := pgx.ConnectConfig(context.Background(), connConfig)
	if err != nil {
		fmt.Errorf("Ошибка при подключении к базе данных:%v", err)
	}
	defer conn.Close(context.Background())

	return conn, nil
}
