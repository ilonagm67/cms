package service

import (
	"cms/internal/usecase"
)

type FSMService struct {
	OrderRepo   usecase.OrderRepository
	ProductRepo usecase.ProductRepository
	UserRepo    usecase.UserRepository
	FSMRepo     usecase.FSMRepository
}

func NewFSMService(FSMRepo usecase.FSMRepository, OrderRepo usecase.OrderRepository, ProductRepo usecase.ProductRepository, UserRepo usecase.UserRepository) *FSMService {
	return &FSMService{OrderRepo: OrderRepo, ProductRepo: ProductRepo, UserRepo: UserRepo, FSMRepo: FSMRepo}
}

func (Service *FSMService) GetCurrentState(id int64) (string, error) {
	state, err := Service.FSMRepo.Get(id)
	if err != nil {
		return "", err
	}
	return string(state), nil
}
