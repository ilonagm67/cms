package telegram

import (
	"cms/internal/service"
	"fmt"

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
	text, err := handler.FSMService.OrderStart(update.Message.Chat.ID)
	if err != nil {
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
			tu.ID(update.Message.Chat.ID),
			fmt.Sprintf(text),
		))
		return err
	}
	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
		tu.ID(update.Message.Chat.ID),
		fmt.Sprintf(text),
	))
	return nil
}

func (handler *OrderHandler) HandleOrderProducts(ctx *th.Context, update telego.Update) error {
	text, err := handler.FSMService.ProcessOrderProducts(update.Message.Chat.ID, update.Message.Text)
	if err != nil {
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
			tu.ID(update.Message.Chat.ID),
			fmt.Sprintf(text),
		))
		return err
	}
	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
		tu.ID(update.Message.Chat.ID),
		fmt.Sprintf(text),
	))
	return nil
}

func (handler *OrderHandler) HandleOrderProductsWeight(ctx *th.Context, update telego.Update) error {
	text, err := handler.FSMService.ProcessOrderProductsWeight(update.Message.Chat.ID, update.Message.Text)
	if err != nil {
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
			tu.ID(update.Message.Chat.ID),
			fmt.Sprintf(text),
		))
		return err
	}
	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
		tu.ID(update.Message.Chat.ID),
		fmt.Sprintf(text),
	))
	return nil
}

func (handler *OrderHandler) HandleOrderProductsCount(ctx *th.Context, update telego.Update) error {
	text, err := handler.FSMService.ProcessOrderProductsCount(update.Message.Chat.ID, update.Message.Text)
	if err != nil {
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
			tu.ID(update.Message.Chat.ID),
			fmt.Sprintf(text),
		))
		return err
	}
	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
		tu.ID(update.Message.Chat.ID),
		fmt.Sprintf(text),
	).WithReplyMarkup(DeliveryKeyboard).WithProtectContent())
	return nil
}

func (handler *OrderHandler) HandleOrderDelivery(ctx *th.Context, update telego.Update) error {
	text, err := handler.FSMService.ProcessOrderDelivery(update.Message.Chat.ID, update.Message.Text)
	if err != nil {
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
			tu.ID(update.Message.Chat.ID),
			fmt.Sprintf(text),
		))
		return err
	}
	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
		tu.ID(update.Message.Chat.ID),
		fmt.Sprintf(text),
	).WithReplyMarkup(PayTypeKeyboard).WithProtectContent())
	return nil
}

func (handler *OrderHandler) HandleOrderPayType(ctx *th.Context, update telego.Update) error {
	text, err := handler.FSMService.ProcessOrderPayType(update.Message.Chat.ID, update.Message.Text)
	if err != nil {
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
			tu.ID(update.Message.Chat.ID),
			fmt.Sprintf(text),
		))
		return err
	}
	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
		tu.ID(update.Message.Chat.ID),
		fmt.Sprintf(text),
	))
	return nil
}

func (handler *OrderHandler) HandleOrderAddress(ctx *th.Context, update telego.Update) error {
	text, err := handler.FSMService.ProcessOrderAddress(update.Message.Chat.ID, update.Message.Text)
	if err != nil {
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
			tu.ID(update.Message.Chat.ID),
			fmt.Sprintf(text),
		))
		return err
	}
	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
		tu.ID(update.Message.Chat.ID),
		fmt.Sprintf(text),
	).WithReplyMarkup(MainKeyboard))
	return nil
}
