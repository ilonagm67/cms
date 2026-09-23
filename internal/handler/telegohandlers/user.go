package telegohandlers

import (
	"cms/internal/service"
	"context"
	"os"
	"strconv"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

type UserHandler struct {
	FSMService  *service.FSMService
	UserService *service.UserService
}

func NewUserHandler(FSMService *service.FSMService, UserService *service.UserService) *UserHandler {
	return &UserHandler{FSMService: FSMService, UserService: UserService}
}

func (handler *UserHandler) StatePredicate(targetState string) th.Predicate {
	return func(ctx context.Context, update telego.Update) bool {
		userID := update.Message.From.ID
		state, err := handler.FSMService.GetCurrentState(userID)
		if err != nil {
			return false
		}
		return state == targetState
	}
}

func (handler *UserHandler) Middleware(ctx *th.Context, update telego.Update) error {
	if update.Message == nil {
		return nil
	}
	if update.Message.Text != "/start" {
		_, err := handler.UserService.Get(update.Message.Chat.ID)
		if err != nil {
			SendMessage(ctx, update.Message.Chat.ID, "Введите /start", nil)
			return err
		}
	}
	return ctx.Next(update)
}

func (handler *UserHandler) HandleQuestionStart(ctx *th.Context, update telego.Update) error {
	err := handler.FSMService.UserQuestionStart(update.Message.Chat.ID)
	if err != nil {
		SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", MainKeyboard)
		return err
	}
	SendMessage(ctx, update.Message.Chat.ID, "Введите вопрос:", nil)
	return nil
}

func (handler *UserHandler) HandleQuestionProcess(ctx *th.Context, update telego.Update) error {
	err := handler.FSMService.UserQuestionProcess(update.Message.Chat.ID)
	if err != nil {
		SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", MainKeyboard)
		return err
	}
	adminID := os.Getenv("ADMIN")
	ID, _ := strconv.ParseInt(adminID, 10, 64)
	SendMessage(ctx, ID, update.Message.Text, nil)
	SendMessage(ctx, update.Message.Chat.ID, "Вопрос отправлен!", MainKeyboard)
	return nil
}

func (handler *UserHandler) HandleStart(ctx *th.Context, update telego.Update) error {
	SendMessage(ctx, update.Message.Chat.ID, "Добро пожаловать в наш магазин!", nil)
	err := handler.FSMService.UserStart(update.Message.Chat.ID)
	if err != nil {
		SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", nil)
		return err
	}
	SendMessage(ctx, update.Message.Chat.ID, "Для регистрации введите ваше имя:", nil)
	return nil
}

func (handler *UserHandler) HandleName(ctx *th.Context, update telego.Update) error {
	err := handler.FSMService.UserProcessName(update.Message.Chat.ID, update.Message.Text)
	if err != nil {
		SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", nil)
		return err
	}
	SendMessage(ctx, update.Message.Chat.ID, "Введите ваш телефон:", NumberKeyboard)
	return nil
}

func (handler *UserHandler) HandlePhone(ctx *th.Context, update telego.Update) error {
	if update.Message.Text == "" {
		err := handler.FSMService.UserProcessPhone(update.Message.Chat.ID, update.Message.Contact.PhoneNumber)
		if err != nil {
			SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", nil)
			return err
		}
	} else {
		err := handler.FSMService.UserProcessPhone(update.Message.Chat.ID, update.Message.Text)
		if err != nil {
			SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", nil)
			return err
		}
	}
	SendMessage(ctx, update.Message.Chat.ID, "Вы успешно зарегистрировались!", nil)
	SendMessage(ctx, update.Message.Chat.ID, "Выберите вариант из списка:", MainKeyboard)
	return nil
}
