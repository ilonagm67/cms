package telegram

import (
	"cms/internal/service"
	"context"
	"fmt"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

type UserHandler struct {
	FSMService *service.FSMService
}

func NewUserHandler(FSMService *service.FSMService) *UserHandler {
	return &UserHandler{FSMService: FSMService}
}

func (handler *UserHandler) StatePredicate(targetState string) th.Predicate {
	return func(ctx context.Context, update telego.Update) bool {
		if update.Message == nil {
			return false
		}
		userID := update.Message.From.ID
		state, err := handler.FSMService.GetCurrentState(userID)
		if err != nil {
			return false
		}
		return state == targetState
	}
}

func (handler *UserHandler) HandleStart(ctx *th.Context, update telego.Update) error {
	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
		tu.ID(update.Message.Chat.ID),
		fmt.Sprintf("Добро пожаловать в наш магазин!"),
	))
	handler.FSMService.UserStart(update.Message.Chat.ID)
	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
		tu.ID(update.Message.Chat.ID),
		fmt.Sprintf("Для регистрации введите ваше имя:"),
	))
	return nil
}

func (handler *UserHandler) HandleName(ctx *th.Context, update telego.Update) error {
	handler.FSMService.UserProcessName(update.Message.Chat.ID, update.Message.Text)
	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
		tu.ID(update.Message.Chat.ID),
		fmt.Sprintf("Введите Ваш Телефон: "),
	).WithReplyMarkup(NumberKeyboard).WithProtectContent())
	return nil
}

func (handler *UserHandler) HandlePhone(ctx *th.Context, update telego.Update) error {
	handler.FSMService.UserProcessPhone(update.Message.Chat.ID, update.Message.Text)
	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
		tu.ID(update.Message.Chat.ID),
		fmt.Sprintf("Вы успешно зарегистрировались!"),
	))
	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
		tu.ID(update.Message.Chat.ID),
		fmt.Sprintf("Выберите вариант из списка:"),
	).WithReplyMarkup(MainKeyboard).WithProtectContent())
	return nil
}
