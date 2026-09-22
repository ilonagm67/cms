package telegohandlers

import (
	"cms/internal/service"
	"errors"
	"strconv"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

type OrderHandler struct {
	FSMService     *service.FSMService
	OrderService   *service.OrderService
	ProductService *service.ProductService
}

func NewOrderHandler(FSMService *service.FSMService, ProductService *service.ProductService, OrderService *service.OrderService) *OrderHandler {
	return &OrderHandler{FSMService: FSMService, ProductService: ProductService, OrderService: OrderService}
}

func (handler *OrderHandler) HandleStart(ctx *th.Context, update telego.Update) error {
	err := handler.FSMService.OrderStart(update.Message.Chat.ID)
	if err != nil {
		handler.FSMService.SetCurrentState(update.Message.Chat.ID, "")
		SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", MainKeyboard)
		return err
	}
	list, _ := handler.ProductService.List()
	for Product := range list {
		err := SendMessage(ctx, update.Message.Chat.ID, Product, nil)
		if err != nil {
			return err
		}
	}
	SendMessage(ctx, update.Message.Chat.ID, "Выберите Продукты:", nil)
	return nil
}

func (handler *OrderHandler) HandleOrderProducts(ctx *th.Context, update telego.Update) error {
	err := handler.FSMService.ProcessOrderProducts(update.Message.Chat.ID, update.Message.Text)
	if err != nil {
		handler.FSMService.SetCurrentState(update.Message.Chat.ID, "")
		SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", MainKeyboard)
		return err
	}
	list, err := handler.ProductService.List()
	if err != nil {
		handler.FSMService.SetCurrentState(update.Message.Chat.ID, "")
		SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", MainKeyboard)
		return err
	}
	for _, Weights := range list {
		for Weight, Product := range Weights {
			if Product.Name == update.Message.Text {
				weight := strconv.Itoa(Weight)
				err := SendMessage(ctx, update.Message.Chat.ID, weight, nil)
				if err != nil {
					return err
				}
			}
		}
	}
	SendMessage(ctx, update.Message.Chat.ID, "Выберите Вес:", nil)
	return nil
}

func (handler *OrderHandler) HandleOrderProductsWeight(ctx *th.Context, update telego.Update) error {
	weight, err := strconv.Atoi(update.Message.Text)
	if err != nil {
		SendMessage(ctx, update.Message.Chat.ID, "Выберите нормальный вес!", nil)
		return err
	}
	err = handler.FSMService.ProcessOrderProductsWeight(update.Message.Chat.ID, weight)
	if err != nil {
		handler.FSMService.SetCurrentState(update.Message.Chat.ID, "")
		SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", MainKeyboard)
		return err
	}
	SendMessage(ctx, update.Message.Chat.ID, "Введите Количество:", nil)
	return nil
}

func (handler *OrderHandler) HandleOrderProductsCount(ctx *th.Context, update telego.Update) error {
	count, err := strconv.Atoi(update.Message.Text)
	if err != nil {
		SendMessage(ctx, update.Message.Chat.ID, "Выберите нормальное количество!", nil)
		return err
	}
	err = handler.FSMService.ProcessOrderProductsCount(update.Message.Chat.ID, count)
	if err != nil {
		handler.FSMService.SetCurrentState(update.Message.Chat.ID, "")
		SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", MainKeyboard)
		return err
	}
	products, err := handler.OrderService.String(update.Message.Chat.ID)
	if err != nil {
		handler.FSMService.SetCurrentState(update.Message.Chat.ID, "")
		SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", MainKeyboard)
		return err
	}
	SendMessage(ctx, update.Message.Chat.ID, products, nil)
	SendMessage(ctx, update.Message.Chat.ID, "Выберите опцию:", PreDeliveryKeyboard)
	return nil
}

func (handler *OrderHandler) HandleOrderProductsDeleteName(ctx *th.Context, update telego.Update) error {
	err := handler.FSMService.ProcessOrderProductsDeleteName(update.Message.Chat.ID, update.Message.Text)
	if err != nil {
		handler.FSMService.SetCurrentState(update.Message.Chat.ID, "")
		SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", MainKeyboard)
		return err
	}
	SendMessage(ctx, update.Message.Chat.ID, "Товар удалён!", nil)
	products, err := handler.OrderService.String(update.Message.Chat.ID)
	if err != nil {
		handler.FSMService.SetCurrentState(update.Message.Chat.ID, "")
		SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", MainKeyboard)
		return err
	}
	SendMessage(ctx, update.Message.Chat.ID, products, nil)
	SendMessage(ctx, update.Message.Chat.ID, "Выберите опцию:", PreDeliveryKeyboard)
	return nil
}

func (handler *OrderHandler) HandleOrderPreDelivery(ctx *th.Context, update telego.Update) error {
	switch update.Message.Text {
	case "Добавить товар":
		err := handler.FSMService.ProcessOrderPreDelivery(update.Message.Chat.ID, update.Message.Text)
		if err != nil {
			handler.FSMService.SetCurrentState(update.Message.Chat.ID, "")
			SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", MainKeyboard)
			return err
		}
		list, _ := handler.ProductService.List()
		for Product := range list {
			err = SendMessage(ctx, update.Message.Chat.ID, Product, nil)
			if err != nil {
				return err
			}
		}
		SendMessage(ctx, update.Message.Chat.ID, "Выберите товар:", nil)
		return nil
	case "Удалить товар":
		err := handler.FSMService.ProcessOrderPreDelivery(update.Message.Chat.ID, update.Message.Text)
		if err != nil {
			handler.FSMService.SetCurrentState(update.Message.Chat.ID, "")
			SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", MainKeyboard)
			return err
		}
		list, _ := handler.ProductService.List()
		for Product := range list {
			err = SendMessage(ctx, update.Message.Chat.ID, Product, nil)
			if err != nil {
				return err
			}
		}
		SendMessage(ctx, update.Message.Chat.ID, "Выберите товар:", nil)
		return nil
	case "Оформить доставку":
		order, _ := handler.OrderService.Get(update.Message.Chat.ID)
		if len(order.Products) == 0 {
			SendMessage(ctx, update.Message.Chat.ID, "Корзина пуста!", nil)
			return errors.New("Cart is empty")
		}
		err := handler.FSMService.ProcessOrderPreDelivery(update.Message.Chat.ID, update.Message.Text)
		if err != nil {
			handler.FSMService.SetCurrentState(update.Message.Chat.ID, "")
			SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", MainKeyboard)
			return err
		}
		SendMessage(ctx, update.Message.Chat.ID, "Выберите опцию:", DeliveryKeyboard)
		return nil
	default:
		SendMessage(ctx, update.Message.Chat.ID, "Выберите опцию!", nil)
		return nil
	}
	return nil
}

func (handler *OrderHandler) HandleOrderDelivery(ctx *th.Context, update telego.Update) error {
	switch update.Message.Text {
	case "Новая Почта":
		err := handler.FSMService.ProcessOrderDelivery(update.Message.Chat.ID, update.Message.Text)
		if err != nil {
			handler.FSMService.SetCurrentState(update.Message.Chat.ID, "")
			SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", MainKeyboard)
			return err
		}
	case "Укр Почта":
		err := handler.FSMService.ProcessOrderDelivery(update.Message.Chat.ID, update.Message.Text)
		if err != nil {
			handler.FSMService.SetCurrentState(update.Message.Chat.ID, "")
			SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", MainKeyboard)
			return err
		}
	case "Самовывоз":
		err := handler.FSMService.ProcessOrderDelivery(update.Message.Chat.ID, update.Message.Text)
		if err != nil {
			handler.FSMService.SetCurrentState(update.Message.Chat.ID, "")
			SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", MainKeyboard)
			return err
		}
	case "Доставка":
		err := handler.FSMService.ProcessOrderDelivery(update.Message.Chat.ID, update.Message.Text)
		if err != nil {
			handler.FSMService.SetCurrentState(update.Message.Chat.ID, "")
			SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", MainKeyboard)
			return err
		}
	default:
		SendMessage(ctx, update.Message.Chat.ID, "Выберите один из вариантов!", nil)
		return errors.New("Delivery Not Found: " + update.Message.Text)
	}
	SendMessage(ctx, update.Message.Chat.ID, "Выберите cпособ оплаты:", PayTypeKeyboard)
	return nil
}

func (handler *OrderHandler) HandleOrderPayType(ctx *th.Context, update telego.Update) error {
	switch update.Message.Text {
	case "Перевод на карту":
		pickup, err := handler.FSMService.ProcessOrderPayType(update.Message.Chat.ID, update.Message.Text)
		if err != nil {
			handler.FSMService.SetCurrentState(update.Message.Chat.ID, "")
			SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", MainKeyboard)
			return err
		}
		if pickup {
			SendMessage(ctx, update.Message.Chat.ID, "Заказ создан!", MainKeyboard)
		} else {
			SendMessage(ctx, update.Message.Chat.ID, "Введите адрес:", nil)
		}
	case "Оплата при получении":
		pickup, err := handler.FSMService.ProcessOrderPayType(update.Message.Chat.ID, update.Message.Text)
		if err != nil {
			handler.FSMService.SetCurrentState(update.Message.Chat.ID, "")
			SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", MainKeyboard)
			return err
		}
		if pickup {
			SendMessage(ctx, update.Message.Chat.ID, "Заказ создан!", MainKeyboard)
		} else {
			SendMessage(ctx, update.Message.Chat.ID, "Введите адрес:", nil)
		}
	default:
		SendMessage(ctx, update.Message.Chat.ID, "Выберите один из вариантов!", nil)
		return errors.New("PayType Not Found: " + update.Message.Text)
	}
	return nil
}

func (handler *OrderHandler) HandleOrderAddress(ctx *th.Context, update telego.Update) error {
	err := handler.FSMService.ProcessOrderAddress(update.Message.Chat.ID, update.Message.Text)
	if err != nil {
		handler.FSMService.SetCurrentState(update.Message.Chat.ID, "")
		SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", MainKeyboard)
		return err
	}
	SendMessage(ctx, update.Message.Chat.ID, "Заказ создан!", MainKeyboard)
	return nil
}
