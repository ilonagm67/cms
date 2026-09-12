package service

import "cms/internal/entity"

func (Service *FSMService) UserQuestionStart(id int64) error {
	return Service.FSMRepo.Set(id, "user_question")
}

func (Service *FSMService) UserQuestionProcess(id int64) error {
	return Service.FSMRepo.Set(id, "")
}

func (Service *FSMService) UserStart(id int64) error {
	err := Service.UserRepo.Add(id, &entity.User{UserID: id})
	if err != nil {
		return err
	}
	return Service.FSMRepo.Set(id, "user_name")
}

func (Service *FSMService) UserProcessName(id int64, name string) error {
	user, err := Service.UserRepo.Get(id)
	if err != nil {
		return err
	}
	user.Name = name
	err = Service.UserRepo.Add(id, user)
	if err != nil {
		return err
	}
	return Service.FSMRepo.Set(id, "user_phone")
}

func (Service *FSMService) UserProcessPhone(id int64, number string) error {
	user, err := Service.UserRepo.Get(id)
	if err != nil {
		return err
	}
	user.Number = number
	err = Service.UserRepo.Add(id, user)
	if err != nil {
		return err
	}
	return Service.FSMRepo.Set(id, "")
}
