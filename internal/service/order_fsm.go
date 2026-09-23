package service

import (
	"cms/internal/entity"
	"errors"
	"fmt"
)

func (Service *FSMService) OrderStart(id int64) error {
	err := Service.OrderRepo.Add(id, &entity.Order{UserID: id})
	if err != nil {
		Service.FSMRepo.Delete(id)
		return err
	}
	_, err = Service.ProductRepo.List()
	if err != nil {
		Service.FSMRepo.Delete(id)
		return err
	}
	Service.FSMRepo.Set(id, "order_products")
	return nil
}

func (Service *FSMService) ProcessOrderProducts(id int64, text string) error {
	err := Service.FSMRepo.SetName(id, text)
	if err != nil {
		Service.FSMRepo.Delete(id)
		return err
	}
	list, err := Service.ProductRepo.List()
	if err != nil {
		Service.FSMRepo.Delete(id)
		return err
	}
	for ProductName := range list {
		if ProductName == text {
			Service.FSMRepo.Set(id, "order_products_weight")
			return nil
		}
	}
	Service.FSMRepo.Delete(id)
	return fmt.Errorf("Product not Found: %s", text)
}

func (Service *FSMService) ProcessOrderProductsWeight(id int64, weight int) error {
	var productmap map[string]map[int]*entity.Product
	err := Service.FSMRepo.SetWeight(id, weight)
	if err != nil {
		Service.FSMRepo.Delete(id)
		return err
	}
	fsmName, err := Service.FSMRepo.GetName(id)
	if err != nil {
		Service.FSMRepo.Delete(id)
		return err
	}
	product, err := Service.ProductRepo.Get(fsmName, weight)
	if err != nil {
		Service.FSMRepo.Delete(id)
		return err
	}
	order, err := Service.OrderRepo.Get(id)
	if err != nil {
		Service.FSMRepo.Delete(id)
		return err
	}
	if order.Products == nil {
		productmap = make(map[string]map[int]*entity.Product, 0)
		_, exists := productmap[product.Name]
		if !exists {
			productmap[product.Name] = make(map[int]*entity.Product)
		}
		productmap[product.Name][product.Weight] = product
		err = Service.OrderRepo.Add(id, &entity.Order{UserID: id, Products: productmap})
		if err != nil {
			Service.FSMRepo.Delete(id)
			return err
		}
	} else {
		productmap = order.Products
		_, exists := productmap[product.Name]
		if !exists {
			productmap[product.Name] = make(map[int]*entity.Product)
		}
		productmap[product.Name][product.Weight] = product
		err = Service.OrderRepo.Add(id, &entity.Order{UserID: id, Products: productmap})
		if err != nil {
			Service.FSMRepo.Delete(id)
			return err
		}
	}
	Service.FSMRepo.Set(id, "order_products_count")
	return nil
}

func (Service *FSMService) ProcessOrderProductsCount(id int64, count int) error {
	fsmName, err := Service.FSMRepo.GetName(id)
	if err != nil {
		Service.FSMRepo.Delete(id)
		return err
	}
	fsmWeight, err := Service.FSMRepo.GetWeight(id)
	if err != nil {
		Service.FSMRepo.Delete(id)
		return err
	}
	order, err := Service.OrderRepo.Get(id)
	if err != nil {
		Service.FSMRepo.Delete(id)
		return err
	}
	order.Products[fsmName][fsmWeight].Count = count
	err = Service.OrderRepo.Add(id, order)
	if err != nil {
		Service.FSMRepo.Delete(id)
		return err
	}
	Service.FSMRepo.Set(id, "order_pre_delivery")
	return nil
}

func (Service *FSMService) ProcessOrderProductsDeleteName(id int64, text string) error {
	err := Service.FSMRepo.SetName(id, text)
	if err != nil {
		Service.FSMRepo.Delete(id)
		return err
	}
	Service.FSMRepo.Set(id, "order_products_delete_weight")
	return nil
}

func (Service *FSMService) ProcessOrderProductsDeleteWeight(id int64, weight int) error {
	order, err := Service.OrderRepo.Get(id)
	if err != nil {
		Service.FSMRepo.Delete(id)
		return err
	}
	fsmName, err := Service.FSMRepo.GetName(id)
	if err != nil {
		Service.FSMRepo.Delete(id)
		return err
	}
	for ProductName, Weights := range order.Products {
		if ProductName == fsmName {
			if len(Weights) > 1 {
				delete(Weights, weight)
			} else {
				delete(order.Products, ProductName)
			}
		}
	}
	Service.FSMRepo.Set(id, "order_pre_delivery")
	return nil
}

func (Service *FSMService) ProcessOrderPreDelivery(id int64, text string) error {
	switch text {
	case "Добавить товар":
		Service.FSMRepo.Set(id, "order_products")
		return nil
	case "Удалить товар":
		Service.FSMRepo.Set(id, "order_products_delete_name")
		return nil
	case "Оформить доставку":
		order, err := Service.OrderRepo.Get(id)
		if err != nil {
			Service.FSMRepo.Delete(id)
			return err
		}
		if len(order.Products) == 0 {
			return errors.New("Cart is empty")
		}
		Service.FSMRepo.Set(id, "order_delivery")
		return nil
	default:
		return errors.New("Text not found!")
	}
	return nil
}

func (Service *FSMService) ProcessOrderDelivery(id int64, text string) error {
	order, err := Service.OrderRepo.Get(id)
	if err != nil {
		Service.FSMRepo.Delete(id)
		return err
	}
	order.Delivery = text
	err = Service.OrderRepo.Add(id, order)
	if err != nil {
		Service.FSMRepo.Delete(id)
		return err
	}
	Service.FSMRepo.Set(id, "order_paytype")
	return nil
}

func (Service *FSMService) ProcessOrderPayType(id int64, text string) (bool, error) {
	order, err := Service.OrderRepo.Get(id)
	if err != nil {
		Service.FSMRepo.Delete(id)
		return false, err
	}
	order.PayType = text
	err = Service.OrderRepo.Add(id, order)
	if err != nil {
		Service.FSMRepo.Delete(id)
		return false, err
	}
	if order.Delivery == "Самовывоз" {
		Service.FSMRepo.Set(id, "order_base_address")
		return true, nil
	}
	Service.FSMRepo.Set(id, "order_address")
	return false, nil
}

func (Service *FSMService) ProcessOrderAddress(id int64, text string) error {
	order, err := Service.OrderRepo.Get(id)
	if err != nil {
		Service.FSMRepo.Delete(id)
		return err
	}
	order.Address = text
	err = Service.OrderRepo.Add(id, order)
	if err != nil {
		Service.FSMRepo.Delete(id)
		return err
	}
	Service.FSMRepo.Delete(id)
	return nil
}
