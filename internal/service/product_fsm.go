package service

import (
	"cms/internal/entity"
	"errors"
	"os"
	"strconv"
)

func (Service *FSMService) ProductCatalogStart(id int64) error {
	err := Service.FSMRepo.Set(id, "product_list")
	if err != nil {
		return err
	}
	return nil
}

func (Service *FSMService) ProductCatalogName(id int64, text string) error {
	adminstr := os.Getenv("ADMIN")
	admin, err := strconv.ParseInt(adminstr, 10, 64)
	if err != nil {
		return err
	}
	if text == "Добавить товар" {
		if admin != id {
			return errors.New("Want Add Product")
		}
		err := Service.FSMRepo.Set(id, "product_name")
		if err != nil {
			return err
		}
		return nil
	} else {
		list, err := Service.ProductRepo.List()
		if err != nil {
			return err
		}
		for Product := range list {
			if text == Product {
				fsm, err := Service.FSMRepo.GetData(id)
				if err != nil {
					return err
				}
				fsm.DataName = text
				fsm.State = "product_list_weight"
				Service.FSMRepo.SetData(id, fsm)
				return nil
			}
		}
	}
	return errors.New("Product not Found!")
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

func (Service *FSMService) ProductAddingWeight(id int64, weight int) (string, error) {
	fsm, err := Service.FSMRepo.GetData(id)
	if err != nil {
		return "Не удалось найти сессию", err
	}
	fsm.DataWeight = weight
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
	product, err := Service.ProductRepo.Get(fsm.DataName, fsm.DataWeight)
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
	product, err := Service.ProductRepo.Get(fsm.DataName, fsm.DataWeight)
	if err != nil {
		return "Не удалось создать продукт!", err
	}
	product.Image = text
	err = Service.ProductRepo.Add(product)
	if err != nil {
		return "Не удалось добавить изображение!", err
	}
	fsm.State = "product_price"
	Service.FSMRepo.SetData(id, fsm)
	return "Введите цену:", nil
}

func (Service *FSMService) ProductAddingPrice(id int64, price int) (string, error) {
	fsm, err := Service.FSMRepo.GetData(id)
	if err != nil {
		return "Не удалось найти сессию", err
	}
	product, err := Service.ProductRepo.Get(fsm.DataName, fsm.DataWeight)
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
