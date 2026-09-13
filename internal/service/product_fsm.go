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
		Service.FSMRepo.SetData(id, text)
		return "Выберите вес:", nil
	}
}

func (Service *FSMService) ProductCatalogWeight(id int64, text string) (*entity.Product, error) {
	weight, err := strconv.Atoi(text)
	if err != nil {
		return nil, err
	}
	name, err := Service.FSMRepo.GetData(id)
	Service.FSMRepo.SetData(id, "")
	if err != nil {
		return nil, err
	}
	product, err := Service.ProductRepo.Get(name, weight)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (Service *FSMService) ProductAddingName(id int64, text string) (string, error) {
	return "", nil
}

func (Service *FSMService) ProductAddingDescription(id int64, text string) (string, error) {
	return "", nil
}

func (Service *FSMService) ProductAddingImage(id int64, text string) (string, error) {
	return "", nil
}

func (Service *FSMService) ProductAddingWeight(id int64, text string) (string, error) {
	return "", nil
}

func (Service *FSMService) ProductAddingPrice(id int64, text string) (string, error) {
	return "", nil
}
