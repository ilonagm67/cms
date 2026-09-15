package memory

import (
	"cms/internal/entity"
	"errors"
)

type UserRepository struct {
	users map[int64]*entity.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[int64]*entity.User, 0),
	}
}

func (repo *UserRepository) Add(ID int64, User *entity.User) error {
	if User.ID == 0 {
		return errors.New("User Not Found")
	}
	repo.users[ID] = User
	return nil
}

func (repo *UserRepository) Delete(ID int64) error {
	_, ok := repo.users[ID]
	if ok {
		delete(repo.users, ID)
		return nil
	}
	return errors.New("User Not Found")
}

func (repo *UserRepository) Get(ID int64) (*entity.User, error) {
	user, ok := repo.users[ID]
	if ok {
		return user, nil
	}
	return nil, errors.New("User Not Found")
}

func (repo *UserRepository) List() (map[int64]*entity.User, error) {
	if len(repo.users) > 0 {
		return repo.users, nil
	}
	return nil, errors.New("No Items in Database Users")
}
