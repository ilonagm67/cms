package service

import "cms/internal/entity"

func (Service *FSMService) OrderStart(id int64) error {
	err := Service.UserRepo.Add(id, &entity.User{UserID: id})
	if err != nil {
		return err
	}
	return Service.FSMRepo.Set(id, "name")
}

func (Service *FSMService) OrderProcessName(id int64, name string) error {
	err := Service.UserRepo.Add(id, &entity.User{UserID: id, Name: name})
	if err != nil {
		return err
	}
	return Service.FSMRepo.Set(id, "phone")
}
