package service

import (
	"cms/internal/entity"
	"cms/internal/usecase"
	"encoding/json"
)

type UserService struct {
	UserRepo usecase.UserRepository
}

func NewUserService(UserRepo usecase.UserRepository) *UserService {
	return &UserService{UserRepo: UserRepo}
}

func (Service *UserService) Add(id int64, name string) error {
	err := Service.UserRepo.Add(id, &entity.User{UserID: id, Name: name})
	if err != nil {
		return err
	}
	return nil
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
	list := Service.UserRepo.List()
	json, err := json.Marshal(list)
	if err != nil {
		return nil, err
	}
	return json, nil
}
