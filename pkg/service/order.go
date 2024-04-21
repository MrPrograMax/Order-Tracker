package service

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"log"
	wbtechl0 "wb-tech-l0"
	"wb-tech-l0/pkg/repository"
)

type OrderService struct {
	repo repository.Order
}

func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(m *stan.Msg) {
	var msg wbtechl0.Order
	if err := json.Unmarshal(m.Data, &msg); err != nil {
		log.Println(err)
		return
	}

	order := wbtechl0.Order{
		OrderUid:          msg.OrderUid,
		TrackNumber:       msg.TrackNumber,
		Entry:             msg.Entry,
		Delivery:          msg.Delivery,
		Payment:           msg.Payment,
		Items:             msg.Items,
		Locale:            msg.Locale,
		InternalSignature: msg.InternalSignature,
		CustomerId:        msg.CustomerId,
		DeliveryService:   msg.DeliveryService,
		ShardKey:          msg.ShardKey,
		SmId:              msg.SmId,
		DateCreated:       msg.DateCreated,
		OofShard:          msg.OofShard,
	}

	err := s.repo.CreateOrder(order)
	if err != nil {
		log.Println(err)
	}
}

func (s *OrderService) GetOrderById(OrderUid string) (wbtechl0.Order, error) {
	order, err := s.repo.GetById(OrderUid)
	return order, err
}
