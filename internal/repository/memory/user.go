package memory

import "cms/internal/entity"

type UserRepository struct {
	users map[int64]*entity.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[int64]*entity.User, 0),
	}
}
