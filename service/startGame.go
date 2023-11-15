package service

import (
	"WB/interal"
	dbCustomer "WB/interal/customer/db"
	"WB/interal/loader"
	dbLoader "WB/interal/loader/db"
	dbTask "WB/interal/task/db"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type gameItem struct {
	ctx  context.Context
	user interal.Model
	sql  *pgxpool.Pool
}

func NewGameItem(ctx context.Context, user interal.Model, sql *pgxpool.Pool) *gameItem {
	return &gameItem{
		ctx:  ctx,
		user: user,
		sql:  sql,
	}
}

func (g *gameItem) StartGame(taskId int, loadersId []int) (int, error) {
	log.Println("Начали игру id task", taskId, loadersId)
	// todo:перписать не стоит всех воркеров беспокоить
	allLoader, err := dbLoader.GetAllLoader(g.ctx, g.sql)
	if err != nil {
		return 2, err
	}
	var weightSum int
	var salarySum int
	mm := make(map[int]loader.Loader)
	for _, ll := range allLoader {
		mm[ll.Id] = ll
	}
	for _, lId := range loadersId {
		salarySum = salarySum + mm[lId].Salary
		weightSum = weightSum + g.portableWeight(mm[lId])
	}
	log.Println(salarySum, weightSum)
	if g.user.Customer.Money < salarySum {
		return 1, fmt.Errorf("Недостаточно денег")
	}

	selectTask, err := dbTask.GetTask(g.ctx, g.sql, taskId)
	if err != nil {
		log.Printf("Get task err :%v", err)
		return 2, err
	}

	if weightSum < selectTask.Weight {
		return 1, fmt.Errorf("Задача оказалась тяжелой")
	}

	g.user.Customer.Money = g.user.Customer.Money - salarySum

	if err := dbTask.UpdateTask(g.ctx, g.sql, taskId); err != nil {
		log.Printf("Update task err :%v", err)
		return 2, err
	}
	if err := dbLoader.UpdateLoader(g.ctx, g.sql, taskId, mm); err != nil {
		log.Printf("Update loader err :%v", err)
		return 2, err
	}
	if err := dbCustomer.UpdateCustomer(g.ctx, g.sql, g.user); err != nil {
		log.Printf("Update Customer err :%v", err)
		return 2, err
	}

	return 0, nil
}

func (g *gameItem) portableWeight(ll loader.Loader) (itog int) {
	itog = ll.Weight * (100 - ll.Tired) / 100
	fmt.Printf("w-%d t-%d i-%d\n-----------\n", ll.Weight, ll.Tired, itog)
	return itog
}
