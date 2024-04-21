package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	wbtechl0 "wb-tech-l0"
)

type OrderPostgres struct {
	db *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

func (r *OrderPostgres) GetById(OrderUid string) (wbtechl0.Order, error) {
	var order wbtechl0.Order
	return order, nil
}

// записал в бд и в кэш
func (r *OrderPostgres) CreateOrder(order wbtechl0.Order) error {
	return nil
}

// обновляем кеш из бд
func (r *OrderPostgres) UploadCache(ctx context.Context) error {
	return nil
}
