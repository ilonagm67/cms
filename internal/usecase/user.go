package usecase

import "cms/internal/entity"

type UserRepository interface {
	Add(ID int64, User *entity.User) error
	SetName(ID int64, Name string) error
	SetPhone(ID int64, Phone string) error
	Delete(ID int64) error
	Get(ID int64) (*entity.User, error)
	List() (map[int64]*entity.User, error)
	Health() error
}
