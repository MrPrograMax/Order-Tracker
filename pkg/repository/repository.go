package repository

import (
	"github.com/jmoiron/sqlx"
	wbtechl0 "wb-tech-l0"
)

type Order interface {
	GetById(OrderUid string) (wbtechl0.Order, error)
	CreateOrder(order wbtechl0.Order) error
	UploadCache() error
}

type Repository struct {
	Order
}

func NewRepository(db *sqlx.DB, cache *Cache) *Repository {
	return &Repository{
		Order: NewOrderPostgres(db, cache),
	}
}
