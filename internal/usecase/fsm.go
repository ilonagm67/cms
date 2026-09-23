package usecase

import "cms/internal/entity"

type FSMRepository interface {
	Set(ID int64, State entity.State)
	SetName(ID int64, Name string) error
	GetName(ID int64) (string, error)
	SetWeight(ID int64, Weight int) error
	GetWeight(ID int64) (int, error)
	Get(ID int64) (entity.State, error)
	Delete(ID int64) error
}
