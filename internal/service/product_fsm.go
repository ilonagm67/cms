package service

import (
	"cms/internal/entity"
	"errors"
	"os"
	"strconv"
)

func (Service *FSMService) ProductCatalogStart(id int64) (bool, error) {
	adminStr := os.Getenv("ADMIN")
	admin, err := strconv.ParseInt(adminStr, 10, 64)
	if err != nil {
		return false, err
	}

	_, err = Service.ProductRepo.List()
	if err != nil {
		if admin != id {
			Service.FSMRepo.Delete(id)
			return false, err
		}
	}
	Service.FSMRepo.Set(id, "product_list")
	if admin == id {
		return true, nil
	} else {
		return false, nil
	}
}

func (Service *FSMService) ProductCatalogName(id int64, text string) error {
	switch text {
	case "Добавить товар":
		Service.FSMRepo.Set(id, "product_name")
		return nil
	default:
		list, err := Service.ProductRepo.List()
		if err != nil {
			Service.FSMRepo.Delete(id)
			return err
		}
		for product, _ := range list {
			if product == text {
				err = Service.FSMRepo.SetName(id, text)
				if err != nil {
					Service.FSMRepo.Delete(id)
					return err
				}
				Service.FSMRepo.Set(id, "product_list_weight")
				return nil
			}
		}
		Service.FSMRepo.Delete(id)
		return errors.New("Product Not Found")
	}
	return nil
}

func (Service *FSMService) ProductCatalogWeight(id int64, weight int) (*entity.Product, error) {
	name, err := Service.FSMRepo.GetName(id)
	if err != nil {
		Service.FSMRepo.Delete(id)
		return nil, err
	}
	product, err := Service.ProductRepo.Get(name, weight)
	if err != nil {
		Service.FSMRepo.Delete(id)
		return nil, err
	}
	Service.FSMRepo.Delete(id)
	return product, nil
}

func (Service *FSMService) ProductAddingName(id int64, text string) error {
	err := Service.FSMRepo.SetName(id, text)
	if err != nil {
		Service.FSMRepo.Delete(id)
		return err
	}
	Service.FSMRepo.Set(id, "product_weight")
	return nil
}

func (Service *FSMService) ProductAddingWeight(id int64, weight int) error {
	name, err := Service.FSMRepo.GetName(id)
	if err != nil {
		Service.FSMRepo.Delete(id)
		return err
	}
	err = Service.FSMRepo.SetWeight(id, weight)
	if err != nil {
		Service.FSMRepo.Delete(id)
		return err
	}
	err = Service.ProductRepo.Add(&entity.Product{Name: name, Weight: weight})
	if err != nil {
		Service.FSMRepo.Delete(id)
		return err
	}
	Service.FSMRepo.Set(id, "product_description")
	return nil
}

func (Service *FSMService) ProductAddingDescription(id int64, text string) error {
	name, err := Service.FSMRepo.GetName(id)
	if err != nil {
		Service.FSMRepo.Delete(id)
		return err
	}
	weight, err := Service.FSMRepo.GetWeight(id)
	if err != nil {
		Service.FSMRepo.Delete(id)
		return err
	}
	product, err := Service.ProductRepo.Get(name, weight)
	if err != nil {
		Service.FSMRepo.Delete(id)
		return err
	}
	product.Description = text
	err = Service.ProductRepo.Add(product)
	if err != nil {
		Service.FSMRepo.Delete(id)
		return err
	}
	Service.FSMRepo.Set(id, "product_image")
	return nil
}

func (Service *FSMService) ProductAddingImage(id int64, text string) error {
	name, err := Service.FSMRepo.GetName(id)
	if err != nil {
		Service.FSMRepo.Delete(id)
		return err
	}
	weight, err := Service.FSMRepo.GetWeight(id)
	if err != nil {
		Service.FSMRepo.Delete(id)
		return err
	}
	product, err := Service.ProductRepo.Get(name, weight)
	if err != nil {
		Service.FSMRepo.Delete(id)
		return err
	}
	product.Image = text
	err = Service.ProductRepo.Add(product)
	if err != nil {
		Service.FSMRepo.Delete(id)
		return err
	}
	Service.FSMRepo.Set(id, "product_price")
	return nil
}

func (Service *FSMService) ProductAddingPrice(id int64, price int) error {
	name, err := Service.FSMRepo.GetName(id)
	if err != nil {
		Service.FSMRepo.Delete(id)
		return err
	}
	weight, err := Service.FSMRepo.GetWeight(id)
	if err != nil {
		Service.FSMRepo.Delete(id)
		return err
	}
	product, err := Service.ProductRepo.Get(name, weight)
	if err != nil {
		Service.FSMRepo.Delete(id)
		return err
	}
	product.Price = price
	err = Service.ProductRepo.Add(product)
	if err != nil {
		Service.FSMRepo.Delete(id)
		return err
	}
	Service.FSMRepo.Delete(id)
	return nil
}
