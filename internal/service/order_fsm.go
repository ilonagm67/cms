package service

import (
	"cms/internal/entity"
)

func (Service *FSMService) OrderStart(id int64) error {
	_, err := Service.OrderRepo.Get(id)
	if err != nil {
		err := Service.OrderRepo.Add(id, &entity.Order{UserID: id})
		if err != nil {
			Service.FSMRepo.Set(id, "")
			return err
		}
		_, err = Service.ProductRepo.List()
		if err != nil {
			Service.FSMRepo.Set(id, "")
			return err
		}
		err = Service.FSMRepo.Set(id, "order_products")
		return nil
	}
	return nil
}

func (Service *FSMService) ProcessOrderProducts(id int64, text string) error {
	fsm, err := Service.FSMRepo.GetData(id)
	if err != nil {
		Service.FSMRepo.Set(id, "")
		return err
	}
	fsm.DataName = text
	fsm.State = "order_products_weight"
	Service.FSMRepo.SetData(id, fsm)
	return nil
}

func (Service *FSMService) ProcessOrderProductsWeight(id int64, weight int) error {
	fsm, err := Service.FSMRepo.GetData(id)
	if err != nil {
		return err
	}
	fsm.DataWeight = weight
	product, _ := Service.ProductRepo.Get(fsm.DataName, fsm.DataWeight)
	productmap := make(map[string]map[int]*entity.Product)
	_, exists := productmap[product.Name]
	if !exists {
		productmap[product.Name] = make(map[int]*entity.Product)
	}
	productmap[product.Name][product.Weight] = product
	err = Service.OrderRepo.Add(id, &entity.Order{UserID: id, Products: productmap})
	if err != nil {
		Service.FSMRepo.Set(id, "")
		return err
	}
	fsm.State = "order_products_count"
	err = Service.FSMRepo.SetData(id, fsm)
	if err != nil {
		Service.FSMRepo.Set(id, "")
		return err
	}
	return nil
}

func (Service *FSMService) ProcessOrderProductsCount(id int64, count int) error {
	fsm, err := Service.FSMRepo.GetData(id)
	if err != nil {
		return err
	}
	order, err := Service.OrderRepo.Get(id)
	if err != nil {
		Service.FSMRepo.Set(id, "")
		return err
	}
	order.Products[fsm.DataName][fsm.DataWeight].Count = count
	err = Service.OrderRepo.Add(id, order)
	if err != nil {
		Service.FSMRepo.Set(id, "")
		return err
	}
	err = Service.FSMRepo.Set(id, "order_pre_delivery")
	if err != nil {
		Service.FSMRepo.Set(id, "")
		return err
	}
	return nil
}

func (Service *FSMService) ProcessOrderPreDelivery(id int64, text string) error {
	switch text {
	case "Добавить товар":
		err := Service.FSMRepo.Set(id, "order_products")
		if err != nil {
			Service.FSMRepo.Set(id, "")
			return err
		}
		return nil
	case "Удалить товар":
	case "Оформить доставку":
	default:
	}
	return nil
}

func (Service *FSMService) ProcessOrderDelivery(id int64, text string) error {
	order, err := Service.OrderRepo.Get(id)
	if err != nil {
		Service.FSMRepo.Set(id, "")
		return err
	}
	order.Delivery = text
	err = Service.OrderRepo.Add(id, order)
	if err != nil {
		Service.FSMRepo.Set(id, "")
		return err
	}
	err = Service.FSMRepo.Set(id, "order_paytype")
	if err != nil {
		Service.FSMRepo.Set(id, "")
		return err
	}
	return nil
}

func (Service *FSMService) ProcessOrderPayType(id int64, text string) (bool, error) {
	order, err := Service.OrderRepo.Get(id)
	if err != nil {
		Service.FSMRepo.Set(id, "")
		return false, err
	}
	order.PayType = text
	err = Service.OrderRepo.Add(id, order)
	if err != nil {
		Service.FSMRepo.Set(id, "")
		return false, err
	}
	if order.Delivery == "Самовывоз" {
		err = Service.FSMRepo.Set(id, "order_base_address")
		if err != nil {
			Service.FSMRepo.Set(id, "")
			return false, err
		}
		return true, nil
	}
	err = Service.FSMRepo.Set(id, "order_address")
	if err != nil {
		Service.FSMRepo.Set(id, "")
		return false, err
	}
	return false, nil
}

func (Service *FSMService) ProcessOrderAddress(id int64, text string) error {
	order, err := Service.OrderRepo.Get(id)
	if err != nil {
		Service.FSMRepo.Set(id, "")
		return err
	}
	order.Address = text
	err = Service.OrderRepo.Add(id, order)
	if err != nil {
		Service.FSMRepo.Set(id, "")
		return err
	}
	err = Service.FSMRepo.Set(id, "")
	if err != nil {
		Service.FSMRepo.Set(id, "")
		return err
	}
	return nil
}
