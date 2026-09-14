package service

import (
	"cms/internal/entity"
)

func (Service *FSMService) OrderStart(id int64) error {
	err := Service.OrderRepo.Add(id, &entity.Order{UserID: id})
	if err != nil {
		return err
	}
	err = Service.FSMRepo.Set(id, "order_products")
	if err != nil {
		return err
	}
	return nil
}

func (Service *FSMService) ProcessOrderProducts(id int64, text string) error {
	order, err := Service.OrderRepo.Get(id)
	if err != nil {
		return err
	}
	order.Products.Name = text
	err = Service.OrderRepo.Add(id, order)
	if err != nil {
		return err
	}
	err = Service.FSMRepo.Set(id, "order_products_weight")
	if err != nil {
		return err
	}
	return nil
}

func (Service *FSMService) ProcessOrderProductsWeight(id int64, weight int) error {
	order, err := Service.OrderRepo.Get(id)
	if err != nil {
		return err
	}
	order.Products.Weight = weight
	err = Service.OrderRepo.Add(id, order)
	if err != nil {
		return err
	}
	Service.FSMRepo.Set(id, "order_products_count")
	if err != nil {
		return err
	}
	return nil
}

func (Service *FSMService) ProcessOrderProductsCount(id int64, count int) error {
	order, err := Service.OrderRepo.Get(id)
	if err != nil {
		return err
	}
	order.Products.Count = count
	err = Service.OrderRepo.Add(id, order)
	if err != nil {
		return err
	}
	err = Service.FSMRepo.Set(id, "order_delivery")
	if err != nil {
		return err
	}
	return nil
}

func (Service *FSMService) ProcessOrderDelivery(id int64, text string) error {
	order, err := Service.OrderRepo.Get(id)
	if err != nil {
		return err
	}
	order.Delivery = text
	err = Service.OrderRepo.Add(id, order)
	if err != nil {
		return err
	}
	err = Service.FSMRepo.Set(id, "order_paytype")
	if err != nil {
		return err
	}
	return nil
}

func (Service *FSMService) ProcessOrderPayType(id int64, text string) (bool, error) {
	order, err := Service.OrderRepo.Get(id)
	if err != nil {
		return false, err
	}
	order.PayType = text
	err = Service.OrderRepo.Add(id, order)
	if err != nil {
		return false, err
	}
	if order.Delivery == "Самовывоз" {
		err = Service.FSMRepo.Set(id, "order_base_address")
		if err != nil {
			return false, err
		}
		return true, nil
	}
	err = Service.FSMRepo.Set(id, "order_address")
	if err != nil {
		return false, err
	}
	return false, nil
}

func (Service *FSMService) ProcessOrderAddress(id int64, text string) error {
	order, err := Service.OrderRepo.Get(id)
	if err != nil {
		return err
	}
	order.Address = text
	err = Service.OrderRepo.Add(id, order)
	if err != nil {
		return err
	}
	err = Service.FSMRepo.Set(id, "")
	if err != nil {
		return err
	}
	return nil
}
