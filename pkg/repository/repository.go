package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	wbtechl0 "wb-tech-l0"
)

type Order interface {
	GetById(OrderUid string) (wbtechl0.Order, error)
	CreateOrder(order wbtechl0.Order) error
	UploadCache(ctx context.Context) error
}

type Repository struct {
	Order
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Order: NewOrderPostgres(db),
	}
}
