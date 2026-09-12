package service

import (
	"cms/internal/entity"
	"errors"
	"fmt"
	"strconv"
)

func (Service *FSMService) OrderStart(id int64) (string, error) {
	err := Service.OrderRepo.Add(id, &entity.Order{UserID: id})
	if err != nil {
		return "Не удалось создать заказ!", err
	}
	Service.FSMRepo.Set(id, "order_products")
	return "Выберите Продукты:", nil
}

func (Service *FSMService) ProcessOrderProducts(id int64, text string) (string, error) {
	order, err := Service.OrderRepo.Get(id)
	if err != nil {
		return "Заказов не найдено!", err
	}
	order.Products.Name = text
	err = Service.OrderRepo.Add(id, order)
	if err != nil {
		return "Не удалось добавить заказ!", err
	}
	Service.FSMRepo.Set(id, "order_products_weight")
	return "Выберите Вес:", nil
}

func (Service *FSMService) ProcessOrderProductsWeight(id int64, text string) (string, error) {
	order, err := Service.OrderRepo.Get(id)
	if err != nil {
		return "Заказов не найдено!", err
	}
	weight, err := strconv.Atoi(text)
	if err != nil {
		return "Выберите нормальный вес!", err
	}
	order.Products.Weight = weight
	err = Service.OrderRepo.Add(id, order)
	if err != nil {
		return "Не удалось добавить заказ!", err
	}
	Service.FSMRepo.Set(id, "order_products_count")
	return "Выберите Количество:", nil
}

func (Service *FSMService) ProcessOrderProductsCount(id int64, text string) (string, error) {
	order, err := Service.OrderRepo.Get(id)
	if err != nil {
		return "Заказов не найдено!", err
	}
	count, err := strconv.Atoi(text)
	if err != nil {
		return "Выберите нормальное количество!", err
	}
	order.Products.Count = count
	err = Service.OrderRepo.Add(id, order)
	if err != nil {
		return "Не удалось добавить заказ!", err
	}
	Service.FSMRepo.Set(id, "order_delivery")
	return "Выберите Доставку:", nil
}

func (Service *FSMService) ProcessOrderDelivery(id int64, text string) (string, error) {
	order, err := Service.OrderRepo.Get(id)
	if err != nil {
		return "Заказов не найдено!", err
	}
	switch text {
	case "Новая Почта":
		order.Delivery = text
	case "Укр Почта":
		order.Delivery = text
	case "Самовывоз":
		order.Delivery = text
	case "Доставка":
		order.Delivery = text
	default:
		return "Выберите один из вариантов!", errors.New(fmt.Sprintf("ID: %v, Enter: %s", id, text))
	}
	err = Service.OrderRepo.Add(id, order)
	if err != nil {
		return "Не удалось добавить заказ!", err
	}
	Service.FSMRepo.Set(id, "order_paytype")
	return "Выберите cпособ оплаты:", nil
}

func (Service *FSMService) ProcessOrderPayType(id int64, text string) (string, error) {
	order, err := Service.OrderRepo.Get(id)
	if err != nil {
		return "Заказов не найдено!", err
	}
	switch text {
	case "Перевод на карту":
		order.PayType = text
	case "Оплата при получении":
		order.PayType = text
	default:
		return "Выберите один из вариантов!", errors.New(fmt.Sprintf("ID: %v, Enter: %s", id, text))
	}
	err = Service.OrderRepo.Add(id, order)
	if err != nil {
		return "Не удалось добавить заказ!", err
	}
	Service.FSMRepo.Set(id, "order_address")
	return "Введите адрес доставки:", nil
}

func (Service *FSMService) ProcessOrderAddress(id int64, text string) (string, error) {
	order, err := Service.OrderRepo.Get(id)
	if err != nil {
		return "Заказов не найдено!", err
	}
	if order.Delivery == "Самовывоз" {
		switch text {
		case "База":
			order.Address = text
		default:
			return "Выберите один из вариантов!", errors.New(fmt.Sprintf("ID: %v, Enter: %s", id, text))
		}
	} else {
		order.Address = text
	}
	err = Service.OrderRepo.Add(id, order)
	if err != nil {
		return "Не удалось добавить заказ!", err
	}
	Service.FSMRepo.Set(id, "")
	return "Заказ создан!", nil
}
