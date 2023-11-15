package db

import (
	"WB/interal"
	"WB/interal/loader"
	"WB/interal/task"
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

func GetInfo(ctx context.Context, conn *pgxpool.Pool, login string) (*loader.Loader, error) {
	q := `
		select weight, money, drunk, tired from loader where login = $1;
	`
	rows := conn.QueryRow(ctx, q, login)
	var load loader.Loader

	if err := rows.Scan(&load.Weight, &load.Salary, &load.Drunk, &load.Tired); err != nil {
		return nil, err
	}

	return &load, nil
}

func GetAllLoader(ctx context.Context, conn *pgxpool.Pool) ([]loader.Loader, error) {
	arrLoader := make([]loader.Loader, 0)
	q := `
		select id, weight, money, drunk, tired from loader;
	`
	rows, err := conn.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var load loader.Loader

		if err := rows.Scan(&load.Id, &load.Weight, &load.Salary, &load.Drunk, &load.Tired); err != nil {
			return nil, err
		}
		arrLoader = append(arrLoader, load)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return arrLoader, nil
}

//todo:check link
func GetTask(ctx context.Context, conn *pgxpool.Pool) ([]task.Task, error) {
	arrTask := make([]task.Task, 0)
	q := `
		select name, weight from task
			inner join loader 
				on task.id = loader.task_id
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

func UpdateLoader(ctx context.Context, conn *pgxpool.Pool) {

}
