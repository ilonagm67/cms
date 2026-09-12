package service

import (
	"cms/internal/entity"
	"strconv"
)

func (Service *FSMService) OrderStart(id int64) error {
	err := Service.OrderRepo.Add(id, &entity.Order{UserID: id})
	if err != nil {
		return err
	}
	return Service.FSMRepo.Set(id, "order_products")
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
	return Service.FSMRepo.Set(id, "order_products_weight")
}

func (Service *FSMService) ProcessOrderProductsWeight(id int64, text string) error {
	order, err := Service.OrderRepo.Get(id)
	if err != nil {
		return err
	}
	weight, err := strconv.Atoi(text)
	if err != nil {
		return err
	}
	order.Products.Weight = weight
	err = Service.OrderRepo.Add(id, order)
	if err != nil {
		return err
	}
	return Service.FSMRepo.Set(id, "order_products_count")
}

func (Service *FSMService) ProcessOrderProductsCount(id int64, text string) error {
	order, err := Service.OrderRepo.Get(id)
	if err != nil {
		return err
	}
	count, err := strconv.Atoi(text)
	if err != nil {
		return err
	}
	order.Products.Count = count
	err = Service.OrderRepo.Add(id, order)
	if err != nil {
		return err
	}
	return Service.FSMRepo.Set(id, "order_delivery")
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
	return Service.FSMRepo.Set(id, "order_paytype")
}

func (Service *FSMService) ProcessOrderPayType(id int64, text string) error {
	order, err := Service.OrderRepo.Get(id)
	if err != nil {
		return err
	}
	order.PayType = text
	err = Service.OrderRepo.Add(id, order)
	if err != nil {
		return err
	}
	return Service.FSMRepo.Set(id, "order_address")
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
	return Service.FSMRepo.Set(id, "")
}
