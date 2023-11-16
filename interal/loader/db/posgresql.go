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

func GetLoaders(ctx context.Context, conn *pgxpool.Pool, id []int) (map[int]loader.Loader, error) {
	ml := make(map[int]loader.Loader, len(id))
	q := `
		select id, weight, money, drunk, tired from loader
			where id = $1
	`
	for _, i := range id {
		rows, err := conn.Query(ctx, q, i)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var load loader.Loader
			if err := rows.Scan(&load.Id, &load.Weight, &load.Salary, &load.Drunk, &load.Tired); err != nil {
				return nil, err
			}
			ml[load.Id] = load
		}
		if err = rows.Err(); err != nil {
			return nil, err
		}
	}
	fmt.Println("map", ml)
	return ml, nil
}

//todo:check link
func GetTask(ctx context.Context, conn *pgxpool.Pool, login string) ([]task.Task, error) {
	var arrTask []task.Task
	fmt.Println("id", login)
	var id int

	q1 := `
		select id from loader
		where login = $1
	`
	row := conn.QueryRow(ctx, q1, login)
	if err := row.Scan(&id); err != nil {
		return nil, err
	}

	q2 := `
        select task.id, task.name, task.weight
        from completed_tasks
        join task on completed_tasks.task_id = task.id
        where completed_tasks.loader_id = $1
    `

	rows, err := conn.Query(ctx, q2, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var data task.Task
		if err := rows.Scan(&data.Id, &data.Name, &data.Weight); err != nil {
			return nil, err
		}
		arrTask = append(arrTask, data)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return arrTask, nil
}

func UpdateLoader(ctx context.Context, conn *pgxpool.Pool, taskId int, loadersId map[int]loader.Loader) (err error) {
	q1 := `
		update loader
			set task_id = $1, tired = $2
			where id = $3
	`
	q2 := `
		insert into completed_tasks (loader_id, task_id)
        values ($1, $2)
	`
	for k, v := range loadersId {
		if v.Drunk {
			v.Tired = +30
		}
		if v.Tired > 80 {
			_, err = conn.Exec(ctx, q1, taskId, 100, k)
		} else {
			_, err = conn.Exec(ctx, q1, taskId, v.Tired+20, k)
		}
		if err != nil {
			return err
		}
		_, err := conn.Exec(ctx, q2, k, taskId)
		if err != nil {
			return err
		}
	}
	return nil
}
