package service

import (
	"cms/internal/entity"
	"strconv"
)

func (Service *FSMService) ProductCatalogStart(id int64) (string, error) {
	Service.FSMRepo.Set(id, "product_list")
	return "Выберите опцию:", nil
}

func (Service *FSMService) ProductCatalogName(id int64, text string) (string, error) {
	switch text {
	case "Добавить товар":
		Service.FSMRepo.Set(id, "product_name")
		return "Введите Название:", nil
	default:
		Service.FSMRepo.Set(id, "product_list_weight")
		fsm, err := Service.FSMRepo.GetData(id)
		if err != nil {
			return "", err
		}
		fsm.DataName = text
		Service.FSMRepo.SetData(id, fsm)
		return "Выберите вес:", nil
	}
}

func (Service *FSMService) ProductCatalogWeight(id int64, weight int) (*entity.Product, error) {
	fsm, err := Service.FSMRepo.GetData(id)
	if err != nil {
		return nil, err
	}
	product, err := Service.ProductRepo.Get(fsm.DataName, weight)
	if err != nil {
		return nil, err
	}
	Service.FSMRepo.Set(id, "")
	return product, nil
}

func (Service *FSMService) ProductAddingName(id int64, text string) (string, error) {
	fsm, err := Service.FSMRepo.GetData(id)
	if err != nil {
		return "", err
	}
	fsm.DataName = text
	fsm.State = "product_weight"
	Service.FSMRepo.SetData(id, fsm)
	return "Введите вес:", nil
}

func (Service *FSMService) ProductAddingWeight(id int64, text string) (string, error) {
	weight, err := strconv.Atoi(text)
	if err != nil {
		return "Введите нормальное число!", err
	}
	fsm, err := Service.FSMRepo.GetData(id)
	if err != nil {
		return "Не удалось найти сессию", err
	}
	fsm.DataWeight = text
	err = Service.ProductRepo.Add(&entity.Product{Name: fsm.DataName, Weight: weight})
	if err != nil {
		return "Не удалось создать продукт!", err
	}
	fsm.State = "product_description"
	Service.FSMRepo.SetData(id, fsm)
	return "Введите Описание:", nil
}

func (Service *FSMService) ProductAddingDescription(id int64, text string) (string, error) {
	fsm, err := Service.FSMRepo.GetData(id)
	if err != nil {
		return "Не удалось найти сессию", err
	}
	weight, _ := strconv.Atoi(fsm.DataWeight)
	product, err := Service.ProductRepo.Get(fsm.DataName, weight)
	if err != nil {
		return "Не удалось создать продукт!", err
	}
	product.Description = text
	err = Service.ProductRepo.Add(product)
	if err != nil {
		return "Не удалось добавить описание!", err
	}
	fsm.State = "product_image"
	Service.FSMRepo.SetData(id, fsm)
	return "Назовите Изображение", nil
}

func (Service *FSMService) ProductAddingImage(id int64, text string) (string, error) {
	fsm, err := Service.FSMRepo.GetData(id)
	if err != nil {
		return "Не удалось найти сессию", err
	}
	weight, _ := strconv.Atoi(fsm.DataWeight)
	product, err := Service.ProductRepo.Get(fsm.DataName, weight)
	if err != nil {
		return "Не удалось создать продукт!", err
	}
	product.Description = text
	err = Service.ProductRepo.Add(product)
	if err != nil {
		return "Не удалось добавить изображение!", err
	}
	fsm.State = "product_price"
	Service.FSMRepo.SetData(id, fsm)
	return "Введите цену:", nil
}

func (Service *FSMService) ProductAddingPrice(id int64, text string) (string, error) {
	fsm, err := Service.FSMRepo.GetData(id)
	if err != nil {
		return "Не удалось найти сессию", err
	}
	weight, _ := strconv.Atoi(fsm.DataWeight)
	price, err := strconv.Atoi(text)
	if err != nil {
		return "Введите нормальную цену!", err
	}
	product, err := Service.ProductRepo.Get(fsm.DataName, weight)
	if err != nil {
		return "Не удалось создать продукт!", err
	}
	product.Price = price
	err = Service.ProductRepo.Add(product)
	if err != nil {
		return "Не удалось добавить цену!", err
	}
	Service.FSMRepo.Set(id, "")
	return "Товар создан!", nil
}
