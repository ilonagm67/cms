package service

import (
	"cms/internal/entity"
	"cms/internal/usecase"
	"encoding/json"
)

type OrderService struct {
	OrderRepo usecase.OrderRepository
}

func NewOrderService(OrderRepo usecase.OrderRepository) *OrderService {
	return &OrderService{OrderRepo: OrderRepo}
}

func (Service *OrderService) Add(id int64, name string) error {
	err := Service.OrderRepo.Add(id, &entity.Order{UserID: id})
	if err != nil {
		return err
	}
	return nil
}

func (Service *OrderService) Get(id int64) ([]byte, error) {
	order, err := Service.OrderRepo.Get(id)
	if err != nil {
		return nil, err
	}
	b, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (Service *OrderService) List() ([]byte, error) {
	list, err := Service.OrderRepo.List()
	if err != nil {
		return nil, err
	}
	json, err := json.Marshal(list)
	if err != nil {
		return nil, err
	}
	return json, nil
}
