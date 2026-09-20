package service

import (
	"cms/internal/entity"
	"cms/internal/usecase"
	"fmt"
)

type OrderService struct {
	OrderRepo usecase.OrderRepository
}

func NewOrderService(OrderRepo usecase.OrderRepository) *OrderService {
	return &OrderService{OrderRepo: OrderRepo}
}

func (Service *OrderService) Get(id int64) (*entity.Order, error) {
	order, err := Service.OrderRepo.Get(id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (Service *OrderService) List() (map[int64]*entity.Order, error) {
	list, err := Service.OrderRepo.List()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (Service *OrderService) String(id int64) (string, error) {
	order, err := Service.OrderRepo.Get(id)
	if err != nil {
		return "", err
	}
	for product, _ := range order.Products {
	}
	return fmt.Sprintf("Ваш заказ:"), nil
}
