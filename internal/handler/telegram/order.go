package telegram

import (
	"cms/internal/entity"
	"cms/internal/service"
	"encoding/json"

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
	text, err := handler.FSMService.OrderStart(update.Message.Chat.ID)
	if err != nil {
		SendMessage(ctx, update.Message.Chat.ID, text, nil)
		return err
	}
	SendMessage(ctx, update.Message.Chat.ID, text, nil)
	return nil
}

func (handler *OrderHandler) HandleOrderProducts(ctx *th.Context, update telego.Update) error {
	text, err := handler.FSMService.ProcessOrderProducts(update.Message.Chat.ID, update.Message.Text)
	if err != nil {
		SendMessage(ctx, update.Message.Chat.ID, text, nil)
		return err
	}
	SendMessage(ctx, update.Message.Chat.ID, text, nil)
	return nil
}

func (handler *OrderHandler) HandleOrderProductsWeight(ctx *th.Context, update telego.Update) error {
	text, err := handler.FSMService.ProcessOrderProductsWeight(update.Message.Chat.ID, update.Message.Text)
	if err != nil {
		SendMessage(ctx, update.Message.Chat.ID, text, nil)
		return err
	}
	SendMessage(ctx, update.Message.Chat.ID, text, nil)
	return nil
}

func (handler *OrderHandler) HandleOrderProductsCount(ctx *th.Context, update telego.Update) error {
	text, err := handler.FSMService.ProcessOrderProductsCount(update.Message.Chat.ID, update.Message.Text)
	if err != nil {
		SendMessage(ctx, update.Message.Chat.ID, text, nil)
		return err
	}
	SendMessage(ctx, update.Message.Chat.ID, text, DeliveryKeyboard)
	return nil
}

func (handler *OrderHandler) HandleOrderDelivery(ctx *th.Context, update telego.Update) error {
	text, err := handler.FSMService.ProcessOrderDelivery(update.Message.Chat.ID, update.Message.Text)
	if err != nil {
		SendMessage(ctx, update.Message.Chat.ID, text, nil)
		return err
	}
	SendMessage(ctx, update.Message.Chat.ID, text, PayTypeKeyboard)
	return nil
}

func (handler *OrderHandler) HandleOrderPayType(ctx *th.Context, update telego.Update) error {
	text, err := handler.FSMService.ProcessOrderPayType(update.Message.Chat.ID, update.Message.Text)
	if err != nil {
		SendMessage(ctx, update.Message.Chat.ID, text, nil)
		return err
	}
	encoded, _ := handler.OrderService.Get(update.Message.Chat.ID)
	var order entity.Order
	err = json.Unmarshal(encoded, &order)
	if err != nil {
		SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", nil)
		return err
	}
	if order.Delivery == "Самовывоз" {
		SendMessage(ctx, update.Message.Chat.ID, text, DeliveryBaseKeyboard)
	} else {
		SendMessage(ctx, update.Message.Chat.ID, text, nil)
	}
	return nil
}

func (handler *OrderHandler) HandleOrderAddress(ctx *th.Context, update telego.Update) error {
	text, err := handler.FSMService.ProcessOrderAddress(update.Message.Chat.ID, update.Message.Text)
	if err != nil {
		SendMessage(ctx, update.Message.Chat.ID, text, nil)
		return err
	}
	SendMessage(ctx, update.Message.Chat.ID, text, MainKeyboard)
	return nil
}
