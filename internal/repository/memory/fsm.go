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

func (repo *FSMRepository) Set(ID int64, State entity.State) error {
	repo.states[ID] = &entity.FSM{UserID: ID, State: State}
	_, ok := repo.states[ID]
	if ok {
		return nil
	}
	return errors.New("FSM not Found")
}

func (repo *FSMRepository) SetData(ID int64, Data string) error {
	State, err := repo.Get(ID)
	if err != nil {
		return errors.New("FSM not Found")
	}
	repo.states[ID] = &entity.FSM{UserID: ID, State: State, Data: Data}
	return nil
}

func (repo *FSMRepository) GetData(ID int64) (string, error) {
	_, ok := repo.states[ID]
	if ok {
		if repo.states[ID].Data == "" {
			return "", errors.New("Data not Found")
		}
		return repo.states[ID].Data, nil
	}
	return "", errors.New("FSM not Found")
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
