package memory

import (
	"cms/internal/entity"
	"errors"
	"fmt"
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

func (repo *OrderRepository) DeleteProduct(ID int64, Name string, Weight int) error {
	products := repo.orders[ID].Products
	weights, productExists := products[Name]
	if !productExists {
		return fmt.Errorf("product category %q not found for customer %d", Name, ID)
	}

	if _, weightExists := weights[Weight]; !weightExists {
		return fmt.Errorf("weight %d not found under product %q", Weight, Name)
	}

	delete(weights, Weight)

	if len(weights) == 0 {
		delete(products, Name)
	}
	return nil
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
	return nil, errors.New("No Items in Database Orders")
}
