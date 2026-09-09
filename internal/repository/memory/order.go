package memory

import (
	"cms/internal/entity"
	"errors"
)

type OrderRepository struct {
	orders      map[int64]map[uint64]*entity.Order
	NextOrderID uint64
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		orders: make(map[int64]map[uint64]*entity.Order, 0),
	}
}

func (repo *OrderRepository) NewOrderID() uint64 {
	repo.NextOrderID++
	return repo.NextOrderID
}

func (repo *OrderRepository) Add(ID int64, Order *entity.Order) error {
	OrderID := repo.NewOrderID()
	_, exists := repo.orders[ID]
	if !exists {
		repo.orders[ID] = make(map[uint64]*entity.Order)
	}
	repo.orders[ID][OrderID] = Order
	return nil
}

func (repo *OrderRepository) Get(ID int64) (*entity.Order, error) {
	_, ok := repo.orders[ID]
	if ok {
		for _, val := range repo.orders[ID] {
			return val, nil
		}
	}
	return nil, errors.New("Order Not Found")
}

func (repo *OrderRepository) List() (map[int64]map[uint64]*entity.Order, error) {
	return repo.orders, nil
}
