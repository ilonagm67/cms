package service

import (
	"cms/internal/usecase"
	"encoding/json"
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

func (Service *UserService) Get(id int64) ([]byte, error) {
	user, err := Service.UserRepo.Get(id)
	if err != nil {
		return nil, err
	}
	b, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (Service *UserService) List() ([]byte, error) {
	list, err := Service.UserRepo.List()
	if err != nil {
		return nil, err
	}
	json, err := json.Marshal(list)
	if err != nil {
		return nil, err
	}
	return json, nil
}

func (Service *UserService) Health() ([]byte, error) {
	message := "Healthy"
	json, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}
	return json, nil
}
