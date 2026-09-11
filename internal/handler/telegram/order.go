package telegram

import (
	"cms/internal/service"
	"fmt"
	"log"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

type OrderHandler struct {
	FSMService *service.FSMService
}

func NewOrderHandler(FSMService *service.FSMService) *OrderHandler {
	return &OrderHandler{FSMService: FSMService}
}

func (handler *OrderHandler) HandleStart(ctx *th.Context, update telego.Update) error {
	err := handler.FSMService.OrderStart(update.Message.Chat.ID)
	if err != nil {
		log.Println(err)
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
			tu.ID(update.Message.Chat.ID),
			fmt.Sprintf("Произошла ошибка!"),
		))
		return err
	}
	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
		tu.ID(update.Message.Chat.ID),
		fmt.Sprintf("Выберите Продукты:"),
	))
	return nil
}

func (handler *OrderHandler) HandleOrderProducts(ctx *th.Context, update telego.Update) error {
	err := handler.FSMService.ProcessOrderProducts(update.Message.Chat.ID, update.Message.Text)
	if err != nil {
		log.Println(err)
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
			tu.ID(update.Message.Chat.ID),
			fmt.Sprintf("Произошла ошибка!"),
		))
		return err
	}
	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
		tu.ID(update.Message.Chat.ID),
		fmt.Sprintf("Выберите Вес:"),
	))
	return nil
}

func (handler *OrderHandler) HandleOrderProductsWeight(ctx *th.Context, update telego.Update) error {
	err := handler.FSMService.ProcessOrderProductsWeight(update.Message.Chat.ID, update.Message.Text)
	if err != nil {
		log.Println(err)
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
			tu.ID(update.Message.Chat.ID),
			fmt.Sprintf("Произошла ошибка!"),
		))
		return err
	}
	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
		tu.ID(update.Message.Chat.ID),
		fmt.Sprintf("Выберите Количество:"),
	))
	return nil
}

func (handler *OrderHandler) HandleOrderProductsCount(ctx *th.Context, update telego.Update) error {
	err := handler.FSMService.ProcessOrderProductsCount(update.Message.Chat.ID, update.Message.Text)
	if err != nil {
		log.Println(err)
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
			tu.ID(update.Message.Chat.ID),
			fmt.Sprintf("Произошла ошибка!"),
		))
		return err
	}
	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
		tu.ID(update.Message.Chat.ID),
		fmt.Sprintf("Выберите Доставку:"),
	).WithReplyMarkup(DeliveryKeyboard).WithProtectContent())
	return nil
}

func (handler *OrderHandler) HandleOrderDelivery(ctx *th.Context, update telego.Update) error {
	err := handler.FSMService.ProcessOrderDelivery(update.Message.Chat.ID, update.Message.Text)
	if err != nil {
		log.Println(err)
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
			tu.ID(update.Message.Chat.ID),
			fmt.Sprintf("Произошла ошибка!"),
		))
		return err
	}
	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
		tu.ID(update.Message.Chat.ID),
		fmt.Sprintf("Выберите cпособ оплаты:"),
	).WithReplyMarkup(PayTypeKeyboard).WithProtectContent())
	return nil
}

func (handler *OrderHandler) HandleOrderPayType(ctx *th.Context, update telego.Update) error {
	err := handler.FSMService.ProcessOrderPayType(update.Message.Chat.ID, update.Message.Text)
	if err != nil {
		log.Println(err)
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
			tu.ID(update.Message.Chat.ID),
			fmt.Sprintf("Произошла ошибка!"),
		))
		return err
	}
	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
		tu.ID(update.Message.Chat.ID),
		fmt.Sprintf("Введите адрес доставки:"),
	))
	return nil
}

func (handler *OrderHandler) HandleOrderAddress(ctx *th.Context, update telego.Update) error {
	err := handler.FSMService.ProcessOrderAddress(update.Message.Chat.ID, update.Message.Text)
	if err != nil {
		log.Println(err)
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
			tu.ID(update.Message.Chat.ID),
			fmt.Sprintf("Произошла ошибка!"),
		))
		return err
	}
	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
		tu.ID(update.Message.Chat.ID),
		fmt.Sprintf("Заказ создан!"),
	).WithReplyMarkup(MainKeyboard))
	return nil
}
