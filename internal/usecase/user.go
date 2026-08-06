package usecase

import "cms/internal/entity"

type  UserRepository interface {
	Add(ID int64,User *entity.User) error
} 
