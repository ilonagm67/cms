package memory

import (
	"cms/internal/entity"
	"fmt"
)

type UserRepository struct {
	users map[int64]*entity.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[int64]*entity.User, 0),
	}
}

func(repo *UserRepository) Add(ID int64,User *entity.User) error {
	repo.users[ID] = User
	fmt.Println(repo.users)
	return nil
}
