package memory

import (
	"cms/internal/entity"
	"errors"
)

type FSMRepository struct {
	states map[int64]*entity.FSM
}

func NewFSMRepository() *FSMRepository {
	return &FSMRepository{
		states: make(map[int64]*entity.FSM, 0),
	}
}

func (repo *FSMRepository) Set(ID int64, State entity.State) {
	_, ok := repo.states[ID]
	if !ok {
		repo.states[ID] = &entity.FSM{UserID: ID, State: State}
	} else {
		repo.states[ID].State = State
	}
}

func (repo *FSMRepository) SetName(ID int64, Name string) error {
	repo.states[ID].DataName = Name
	return nil
}

func (repo *FSMRepository) GetName(ID int64) (string, error) {
	_, ok := repo.states[ID]
	if !ok {
		return "", errors.New("Data Not Found")
	}
	if repo.states[ID].DataName == "" {
		return "", errors.New("DataName is empty")
	}
	return repo.states[ID].DataName, nil
}

func (repo *FSMRepository) SetWeight(ID int64, Weight int) error {
	repo.states[ID].DataWeight = Weight
	return nil
}

func (repo *FSMRepository) GetWeight(ID int64) (int, error) {
	_, ok := repo.states[ID]
	if !ok {
		return 0, errors.New("Data Not Found")
	}
	if repo.states[ID].DataWeight == 0 {
		return 0, errors.New("DataWeight is empty")
	}
	return repo.states[ID].DataWeight, nil
}

func (repo *FSMRepository) Delete(ID int64) error {
	_, ok := repo.states[ID]
	if ok {
		delete(repo.states, ID)
		return nil
	}
	return errors.New("FSM not Found")
}

func (repo *FSMRepository) Get(ID int64) (entity.State, error) {
	fsm, ok := repo.states[ID]
	if ok {
		return fsm.State, nil
	}
	return "", errors.New("FSM not Found")
}
