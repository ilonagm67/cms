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
	repo.users[ID] = User
	repo.users[ID].State = "none"
	return nil
}

func (repo *UserRepository) SetState(ID int64, state entity.State) error {
	repo.users[ID].State = state
	return nil
}

func (repo *UserRepository) GetState(ID int64) (entity.State, error) {
	return repo.users[ID].State, nil
}

func (repo *UserRepository) Delete(ID int64) error {
	_, ok := repo.users[ID]
	if ok {
		delete(repo.users, ID)
		return nil
	} else {
		return errors.New("User Not Found")
	}
}

func (repo *UserRepository) Get(ID int64) (*entity.User, error) {
	i, ok := repo.users[ID]
	if ok {
		return i, nil
	} else {
		return nil, errors.New("User Not Found")
	}
}

func (repo *UserRepository) List() map[int64]*entity.User {
	return repo.users
}
