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
	var cart string
	var Total int
	order, err := Service.OrderRepo.Get(id)
	if err != nil {
		return "", err
	}
	count := 1
	for _, weights := range order.Products {
		for _, product := range weights {
			ProductTotal := product.Price * product.Count
			cart += fmt.Sprintf("%d,Название: %s\nВес: %d\nКоличество: %d\nЦена за ед: %d\nОбщая Цена: %d\n\n", count, product.Name, product.Weight, product.Count, product.Price, ProductTotal)
			Total += product.Price * product.Count
			count++
		}
	}
	return fmt.Sprintf("Ваш заказ:\n\n%s\nСумма: %d", cart, Total), nil
}
