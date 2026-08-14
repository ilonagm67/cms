package memory

import (
	"cms/internal/entity"
	"errors"
)

type OrderRepository struct {
	orders map[int64]*entity.Order
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		orders: make(map[int64]*entity.Order, 0),
	}
}

func (repo *OrderRepository) Add(ID int64, Order *entity.Order) error {
	repo.orders[ID] = Order
	return nil
}

func (repo *OrderRepository) Get(ID int64) (*entity.Order, error) {
	i, ok := repo.orders[ID]
	if ok {
		return i, nil
	} else {
		return nil, errors.New("Order Not Found")
	}
}

func (repo *OrderRepository) List() map[int64]*entity.Order {
	return repo.orders
}
