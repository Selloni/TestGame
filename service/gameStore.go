package service

import (
	"WB/interal"
	"WB/interal/loader"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
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

type GameStoreI interface {
	StartGame(taskId int, loadersId []int) (int, error)
	portableWeight(ll loader.Loader) (itog int)
}
