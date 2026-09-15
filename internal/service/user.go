package service

import (
	"cms/internal/entity"
	"cms/internal/usecase"
)

type UserService struct {
	UserRepo usecase.UserRepository
}

func NewUserService(UserRepo usecase.UserRepository) *UserService {
	return &UserService{UserRepo: UserRepo}
}

func (Service *UserService) Delete(id int64) error {
	err := Service.UserRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (Service *UserService) Get(id int64) (*entity.User, error) {
	user, err := Service.UserRepo.Get(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (Service *UserService) List() (map[int64]*entity.User, error) {
	list, err := Service.UserRepo.List()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (Service *UserService) Health() (string, error) {
	message := "Healthy"
	return message, nil
}
