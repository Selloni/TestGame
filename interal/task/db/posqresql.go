package db

import (
	"WB/interal/task"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"math/rand"
)

func CreateTask(ctx context.Context, conn *pgxpool.Pool) (map[string]int, error) {
	mm := generateItem()
	q := `
		insert into task
			(name, weight)
		values
		    ($1,$2)
	`
	for str, i := range mm {
		err := conn.QueryRow(ctx, q, str, i).Scan()
		if err != nil {
			if err == pgx.ErrNoRows {
				continue
			}
			return nil, fmt.Errorf("не удалось создать груз")
		}
	}
	return mm, nil
}

func UpdateTask(ctx context.Context, conn *pgxpool.Pool, id int) error {
	q := `
		update task
			set done = true
			where id = $1
	`
	_, err := conn.Exec(ctx, q, id)
	if err != nil {
		return err
	}
	return nil
}

func GetTask(ctx context.Context, conn *pgxpool.Pool, item int) (oneTask task.Task, err error) {
	q := `
			select name, weight from task where id = $1;
		`
	rows := conn.QueryRow(ctx, q, item)

	if err = rows.Scan(&oneTask.Name, &oneTask.Weight); err != nil {
		return task.Task{}, err
	}
	return oneTask, nil
}

// support func
func generateItem() map[string]int {
	countItem := rand.Intn(4) + 2

	mm := make(map[string]int, countItem)
	//rand.Seed(time.Now().Unix()) // портит
	length := 4
	ranStr := make([]byte, length)

	for ; countItem > 0; countItem-- {
		for i := 0; i < length; i++ {
			ranStr[i] = byte(65 + rand.Intn(25))
		}
		mm[string(ranStr)] = rand.Intn(70) + 10
	}
	return mm
}
