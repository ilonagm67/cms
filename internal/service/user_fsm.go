package service

import "cms/internal/entity"

func (Service *FSMService) UserQuestionStart(id int64) error {
	Service.FSMRepo.Set(id, "user_question")
	return nil
}

func (Service *FSMService) UserQuestionProcess(id int64) error {
	err := Service.FSMRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (Service *FSMService) UserStart(id int64) error {
	err := Service.UserRepo.Add(id, &entity.User{ID: id})
	if err != nil {
		return err
	}
	Service.FSMRepo.Set(id, "user_name")
	return nil
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
	Service.FSMRepo.Set(id, "user_phone")
	return nil
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
	err = Service.FSMRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
