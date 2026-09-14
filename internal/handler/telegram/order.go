package telegram

import (
	"cms/internal/service"
	"errors"
	"strconv"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

type OrderHandler struct {
	FSMService   *service.FSMService
	OrderService *service.OrderService
}

func NewOrderHandler(FSMService *service.FSMService, OrderService *service.OrderService) *OrderHandler {
	return &OrderHandler{FSMService: FSMService, OrderService: OrderService}
}

func (handler *OrderHandler) HandleStart(ctx *th.Context, update telego.Update) error {
	err := handler.FSMService.OrderStart(update.Message.Chat.ID)
	if err != nil {
		SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", nil)
		return err
	}
	SendMessage(ctx, update.Message.Chat.ID, "Выберите Продукты:", nil)
	return nil
}

func (handler *OrderHandler) HandleOrderProducts(ctx *th.Context, update telego.Update) error {
	err := handler.FSMService.ProcessOrderProducts(update.Message.Chat.ID, update.Message.Text)
	if err != nil {
		SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", nil)
		return err
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
		SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", nil)
		return err
	}
	SendMessage(ctx, update.Message.Chat.ID, "Выберите Количество:", nil)
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
		SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", nil)
		return err
	}
	SendMessage(ctx, update.Message.Chat.ID, "Выберите Доставку:", DeliveryKeyboard)
	return nil
}

func (handler *OrderHandler) HandleOrderDelivery(ctx *th.Context, update telego.Update) error {
	switch update.Message.Text {
	case "Новая Почта":
		err := handler.FSMService.ProcessOrderDelivery(update.Message.Chat.ID, update.Message.Text)
		if err != nil {
			SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", nil)
			return err
		}
	case "Укр Почта":
		err := handler.FSMService.ProcessOrderDelivery(update.Message.Chat.ID, update.Message.Text)
		if err != nil {
			SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", nil)
			return err
		}
	case "Самовывоз":
		err := handler.FSMService.ProcessOrderDelivery(update.Message.Chat.ID, update.Message.Text)
		if err != nil {
			SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", nil)
			return err
		}
	case "Доставка":
		err := handler.FSMService.ProcessOrderDelivery(update.Message.Chat.ID, update.Message.Text)
		if err != nil {
			SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", nil)
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
			SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", nil)
			return err
		}
		if pickup {
			SendMessage(ctx, update.Message.Chat.ID, "Заказ создан!", MainKeyboard)
		} else {
			SendMessage(ctx, update.Message.Chat.ID, "", nil)
		}
	case "Оплата при получении":
		pickup, err := handler.FSMService.ProcessOrderPayType(update.Message.Chat.ID, update.Message.Text)
		if err != nil {
			SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", nil)
			return err
		}
		if pickup {
			SendMessage(ctx, update.Message.Chat.ID, "Заказ создан!", MainKeyboard)
		} else {
			SendMessage(ctx, update.Message.Chat.ID, "", nil)
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
		SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", MainKeyboard)
		return err
	}
	SendMessage(ctx, update.Message.Chat.ID, "Заказ создан!", MainKeyboard)
	return nil
}
