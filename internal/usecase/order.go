package usecase

import "cms/internal/entity"

type OrderRepository interface {
	Add(ID int64, Order *entity.Order) error
	Get(ID int64) (*entity.Order, error)
	List() (map[int64]*entity.Order, error)
}
