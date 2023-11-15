package service

import (
	"WB/interal"
	"WB/interal/loader"
	"WB/interal/loader/db"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type GameItem struct {
	user interal.Model
	ctx  context.Context
	sql  *pgxpool.Pool
}

func (g *GameItem) StartGame(taskId int, loadersId []int) (bool, error) {
	allLoader, err := db.GetAllLoader(g.ctx, g.sql)
	var weightSum int
	var salarySum int
	mm := make(map[int]loader.Loader)
	if err != nil {
		return false, err
	}
	for _, ll := range allLoader {
		mm[ll.Id] = ll
	}
	for _, lid := range loadersId {
		g.PortableWeight(mm[lid])
		salarySum = salarySum + mm[lid].Salary
		weightSum = weightSum + g.PortableWeight(mm[lid])
	}
	return false, nil
}

func (g *GameItem) PortableWeight(ll loader.Loader) (itog int) {
	if ll.Drunk {
		itog = ll.Weight * (100 - (ll.Tired+50)/100) * (2 / 100)
	} else {
		itog = ll.Weight * (100 - ll.Tired/100) * (1 / 100)
	}
	return
}
