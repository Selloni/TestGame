package db

import (
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

func UpdateItem(ctx context.Context, conn *pgxpool.Pool, tt []int) error {
	q := `
		update loader
			set done = true
			where id = &1
	`
	for _, id := range tt {
		_, err := conn.Exec(ctx, q, id)
		if err != nil {
			return err
		}
	}
	return nil
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
