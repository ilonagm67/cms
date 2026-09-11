package service

import "cms/internal/usecase"

type FSMService struct {
	FSMRepo usecase.FSMRepository
}

func NewFSMService(FSMRepo usecase.FSMRepository) *FSMService {
	return &FSMService{FSMRepo: FSMRepo}
}

func (Service *FSMService) GetCurrentState(id int64) (string, error) {
	state, err := Service.FSMRepo.Get(id)
	if err != nil {
		return "", err
	}
	return string(state), nil
}
