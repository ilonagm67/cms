package service

import "cms/internal/entity"

func (Service *FSMService) UserQuestionStart(id int64) (string, error) {
	Service.FSMRepo.Set(id, "user_question")
	return "Введите вопрос:", nil
}

func (Service *FSMService) UserQuestionProcess(id int64) (string, error) {
	Service.FSMRepo.Set(id, "")
	return "Вопрос отправлен!", nil
}

func (Service *FSMService) UserStart(id int64) (string, error) {
	err := Service.UserRepo.Add(id, &entity.User{UserID: id})
	if err != nil {
		return "Не удалось добавить пользователя!", err
	}
	Service.FSMRepo.Set(id, "user_name")
	return "Для регистрации введите ваше имя:", nil
}

func (Service *FSMService) UserProcessName(id int64, name string) (string, error) {
	user, err := Service.UserRepo.Get(id)
	if err != nil {
		return "Не удалось найти пользователя!", err
	}
	user.Name = name
	err = Service.UserRepo.Add(id, user)
	if err != nil {
		return "Не удалось добавить пользователя!", err
	}
	Service.FSMRepo.Set(id, "user_phone")
	return "Введите ваш телефон:", nil
}

func (Service *FSMService) UserProcessPhone(id int64, number string) (string, error) {
	user, err := Service.UserRepo.Get(id)
	if err != nil {
		return "Не удалось найти пользователя!", err
	}
	user.Number = number
	err = Service.UserRepo.Add(id, user)
	if err != nil {
		return "Не удалось добавить пользователя!", err
	}
	Service.FSMRepo.Set(id, "")
	return "Вы успешно зарегистрировались!", nil
}
