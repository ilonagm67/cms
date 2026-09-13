package usecase

import "cms/internal/entity"

type FSMRepository interface {
	Set(ID int64, State entity.State) error
	SetData(ID int64, Data string) error
	GetData(ID int64) (string, error)
	Get(ID int64) (entity.State, error)
	Delete(ID int64) error
}
