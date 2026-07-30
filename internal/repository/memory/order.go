package memory

import "cms/internal/entity"

type OrderRepository struct {
	orders map[int64]*entity.Order
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		orders: make(map[int64]*entity.Order, 0),
	}
}
