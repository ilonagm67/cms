package service

import (
	"cms/internal/usecase"
	"cms/internal/entity"
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
