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
	_, ok := repo.orders[ID]
	if ok {
		return nil
	}
	return errors.New("Order Not Created")
}

func (repo *OrderRepository) Get(ID int64) (*entity.Order, error) {
	_, ok := repo.orders[ID]
	if ok {
		return repo.orders[ID], nil
	}
	return nil, errors.New("Order Not Found")
}

func (repo *OrderRepository) List() (map[int64]*entity.Order, error) {
	if len(repo.orders) > 0 {
		return repo.orders, nil
	}
	return nil, errors.New("No Items in Database")
}
