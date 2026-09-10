package usecase

import "cms/internal/entity"

type FSMRepository interface {
	Set(ID int64, State entity.State) error
	Get(ID int64) (entity.State, error)
	Delete(ID int64) error
}
