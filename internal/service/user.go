package service

import (
	"cms/internal/usecase"
	"cms/internal/entity"
	"encoding/json"
)

type UserService struct {
	UserRepo usecase.UserRepository
}

func NewUserService(UserRepo usecase.UserRepository) *UserService {
	return &UserService{UserRepo: UserRepo}
}

func (Service *UserService) Add(id int64,name string) {
	Service.UserRepo.Add(id,&entity.User{ID:id,Name: name})
}

func (Service *UserService) Delete(id int64) bool {
	err := Service.UserRepo.Delete(id)
	if err != nil {
		return false
	}
	return true
}

func (Service *UserService) Get(id int64) ([]byte,error) {
	user,err := Service.UserRepo.Get(id)
	if err != nil {
		return nil,err
	}
	b,err := json.Marshal(user)
	if err != nil {
		return nil,err
	}
	return b,nil
}

func (Service *UserService) List() ([]byte,error) {
	list := Service.UserRepo.List()
	json,err := json.Marshal(list)
	if err != nil {
		return nil,err
	}
	return json,nil
}
