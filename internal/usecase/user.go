package usecase

import "cms/internal/entity"

type UserRepository interface {
	Add(User *entity.User) error
	Delete(ID int64) error
	Get(ID int64) (*entity.User, error)
	List() (map[int64]*entity.User, error)
}
