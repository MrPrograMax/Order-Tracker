package service

import (
	"github.com/nats-io/stan.go"
	wbtechl0 "wb-tech-l0"
	"wb-tech-l0/pkg/repository"
)

type Order interface {
	CreateOrder(m *stan.Msg)
	GetOrderById(OrderUid string) (wbtechl0.Order, error)
}

type Service struct {
	Order
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Order: NewOrderService(repos.Order),
	}
}
