package service

import (
	"cms/internal/entity"
	"errors"
)

func (Service *FSMService) ProductCatalogStart(id int64, admin bool) error {
	_, err := Service.ProductRepo.List()
	if err != nil {
		if !admin {
			Service.FSMRepo.Set(id, "")
		}
	}
	Service.FSMRepo.Set(id, "product_list")
	return nil
}

func (Service *FSMService) ProductCatalogName(id int64, text string) error {
	switch text {
	case "Добавить товар":
		Service.FSMRepo.Set(id, "product_name")
		return nil
	default:
		fsm, err := Service.FSMRepo.GetData(id)
		if err != nil {
			Service.FSMRepo.Set(id, "")
			return err
		}
		list, err := Service.ProductRepo.List()
		if err != nil {
			Service.FSMRepo.Set(id, "")
			return err
		}
		for product, _ := range list {
			if product == text {
				fsm.DataName = text
				fsm.State = "product_list_weight"
				Service.FSMRepo.SetData(id, fsm)
				return nil
			}
		}
		Service.FSMRepo.Set(id, "")
		return errors.New("Product Not Found")
	}
	return nil
}

func (Service *FSMService) ProductCatalogWeight(id int64, weight int) (*entity.Product, error) {
	fsm, err := Service.FSMRepo.GetData(id)
	if err != nil {
		Service.FSMRepo.Set(id, "")
		return nil, err
	}
	product, err := Service.ProductRepo.Get(fsm.DataName, weight)
	if err != nil {
		Service.FSMRepo.Set(id, "")
		return nil, err
	}
	Service.FSMRepo.Set(id, "")
	return product, nil
}

func (Service *FSMService) ProductAddingName(id int64, text string) error {
	fsm, err := Service.FSMRepo.GetData(id)
	if err != nil {
		Service.FSMRepo.Set(id, "")
		return err
	}
	fsm.DataName = text
	fsm.State = "product_weight"
	Service.FSMRepo.SetData(id, fsm)
	return nil
}

func (Service *FSMService) ProductAddingWeight(id int64, weight int) error {
	fsm, err := Service.FSMRepo.GetData(id)
	if err != nil {
		return err
	}
	fsm.DataWeight = weight
	err = Service.ProductRepo.Add(&entity.Product{Name: fsm.DataName, Weight: weight})
	if err != nil {
		return err
	}
	fsm.State = "product_description"
	Service.FSMRepo.SetData(id, fsm)
	return nil
}

func (Service *FSMService) ProductAddingDescription(id int64, text string) error {
	fsm, err := Service.FSMRepo.GetData(id)
	if err != nil {
		return err
	}
	product, err := Service.ProductRepo.Get(fsm.DataName, fsm.DataWeight)
	if err != nil {
		return err
	}
	product.Description = text
	err = Service.ProductRepo.Add(product)
	if err != nil {
		return err
	}
	fsm.State = "product_image"
	Service.FSMRepo.SetData(id, fsm)
	return nil
}

func (Service *FSMService) ProductAddingImage(id int64, text string) error {
	fsm, err := Service.FSMRepo.GetData(id)
	if err != nil {
		return err
	}
	product, err := Service.ProductRepo.Get(fsm.DataName, fsm.DataWeight)
	if err != nil {
		return err
	}
	product.Image = text
	err = Service.ProductRepo.Add(product)
	if err != nil {
		return err
	}
	fsm.State = "product_price"
	Service.FSMRepo.SetData(id, fsm)
	return nil
}

func (Service *FSMService) ProductAddingPrice(id int64, price int) error {
	fsm, err := Service.FSMRepo.GetData(id)
	if err != nil {
		return err
	}
	product, err := Service.ProductRepo.Get(fsm.DataName, fsm.DataWeight)
	if err != nil {
		return err
	}
	product.Price = price
	err = Service.ProductRepo.Add(product)
	if err != nil {
		return err
	}
	Service.FSMRepo.Set(id, "")
	return nil
}
