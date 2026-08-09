package usecase

import "cms/internal/entity"

type OrderRepository interface {
	Add(ID uint64,Order *entity.Order) error
	Get(ID uint64) (*entity.Order,error)
	List() (map[uint64]*entity.Order)
} 
