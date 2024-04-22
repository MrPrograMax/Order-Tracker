package repository

import (
	"fmt"
	wbtechl0 "wb-tech-l0"
)

type Cache struct {
	cache map[string]wbtechl0.Order
}

func NewCache() *Cache {
	return &Cache{cache: make(map[string]wbtechl0.Order)}
}

func (c *Cache) GetOrder(OrderUid string) (wbtechl0.Order, error) {
	if order, status := c.cache[OrderUid]; status {
		return order, nil
	}

	return wbtechl0.Order{}, fmt.Errorf("order not found")
}

func (c *Cache) AddOrder(order wbtechl0.Order) {
	c.cache[order.OrderUid] = order
}
