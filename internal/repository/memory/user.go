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

func (repo *UserRepository) Add(ID int64, User *entity.User) error {
	if User.ID == 0 {
		return fmt.Errorf("User ID Not Found")
	}
	repo.users[ID] = User
	return nil
}

func (repo *UserRepository) SetName(ID int64, Name string) error {
	if ID == 0 || Name == "" {
		return fmt.Errorf("ID or Name not set!")
	}
	_, ok := repo.users[ID]
	if !ok {
		return fmt.Errorf("User ID: %d Not Found!", ID)
	}
	repo.users[ID].Name = Name
	return nil
}

func (repo *UserRepository) SetPhone(ID int64, Phone string) error {
	if ID == 0 || Phone == "" {
		return fmt.Errorf("ID or Phone not set!")
	}
	_, ok := repo.users[ID]
	if !ok {
		return fmt.Errorf("User ID: %d Not Found!", ID)
	}
	repo.users[ID].Number = Phone
	return nil
}

func (repo *UserRepository) Delete(ID int64) error {
	_, ok := repo.users[ID]
	if !ok {
		return fmt.Errorf("User ID: %d Not Found", ID)
	}
	delete(repo.users, ID)
	return nil
}

func (repo *UserRepository) Get(ID int64) (*entity.User, error) {
	user, ok := repo.users[ID]
	if !ok {
		return nil, fmt.Errorf("User ID: %d Not Found", ID)
	}
	return user, nil
}

func (repo *UserRepository) List() (map[int64]*entity.User, error) {
	if len(repo.users) == 0 {
		return nil, fmt.Errorf("No Items in Database Users")
	}
	return repo.users, nil
}

func (repo *UserRepository) Health() error {
	return nil
}
