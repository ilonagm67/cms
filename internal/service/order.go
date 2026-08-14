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

func (Service *OrderService) Add(id int64, name string) {
	Service.OrderRepo.Add(id, &entity.Order{ID: 1})
}

func (Service *OrderService) Get(id int64) ([]byte, error) {
	user, err := Service.OrderRepo.Get(id)
	if err != nil {
		return nil, err
	}
	b, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (Service *OrderService) List() ([]byte, error) {
	list := Service.OrderRepo.List()
	json, err := json.Marshal(list)
	if err != nil {
		return nil, err
	}
	return json, nil
}
