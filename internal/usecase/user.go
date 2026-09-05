package usecase

import "cms/internal/entity"

type UserRepository interface {
	Add(ID int64, User *entity.User) error
	SetState(ID int64, state entity.State) error
	GetState(ID int64) (entity.State, error)
	Delete(ID int64) error
	Get(ID int64) (*entity.User, error)
	List() map[int64]*entity.User
}
