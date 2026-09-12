package telegram

import (
	"cms/internal/service"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
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

func (handler *UserHandler) Middleware(ctx *th.Context, update telego.Update) error {
	if update.Message.Text != "/start" {
		_, err := handler.UserService.Get(update.Message.Chat.ID)
		if err != nil {
			_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
				tu.ID(update.Message.Chat.ID),
				fmt.Sprintf("Введите /start"),
			))
			return err
		}
	}
	return ctx.Next(update)
}

func (handler *UserHandler) HandleQuestionStart(ctx *th.Context, update telego.Update) error {
	err := handler.FSMService.UserQuestionStart(update.Message.Chat.ID)
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
		fmt.Sprintf("Введите вопрос:"),
	))
	return err
}

func (handler *UserHandler) HandleQuestionProcess(ctx *th.Context, update telego.Update) error {
	err := handler.FSMService.UserQuestionProcess(update.Message.Chat.ID)
	if err != nil {
		log.Println(err)
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
			tu.ID(update.Message.Chat.ID),
			fmt.Sprintf("Произошла ошибка!"),
		))
		return err
	}
	adminID := os.Getenv("ADMIN")
	ID, _ := strconv.ParseInt(adminID, 10, 64)
	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
		tu.ID(ID),
		fmt.Sprintf(update.Message.Text),
	))
	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
		tu.ID(update.Message.Chat.ID),
		fmt.Sprintf("Вопрос отправлен!"),
	))
	return nil
}

func (handler *UserHandler) HandleStart(ctx *th.Context, update telego.Update) error {
	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
		tu.ID(update.Message.Chat.ID),
		fmt.Sprintf("Добро пожаловать в наш магазин!"),
	))
	err := handler.FSMService.UserStart(update.Message.Chat.ID)
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
		fmt.Sprintf("Для регистрации введите ваше имя:"),
	))
	return nil
}

func (handler *UserHandler) HandleName(ctx *th.Context, update telego.Update) error {
	err := handler.FSMService.UserProcessName(update.Message.Chat.ID, update.Message.Text)
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
		fmt.Sprintf("Введите Ваш Телефон:"),
	).WithReplyMarkup(NumberKeyboard).WithProtectContent())
	return nil
}

func (handler *UserHandler) HandlePhone(ctx *th.Context, update telego.Update) error {
	if update.Message.Text == "" {
		err := handler.FSMService.UserProcessPhone(update.Message.Chat.ID, update.Message.Contact.PhoneNumber)
		if err != nil {
			log.Println(err)
			_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
				tu.ID(update.Message.Chat.ID),
				fmt.Sprintf("Произошла ошибка!"),
			))
			return err
		}
	} else {
		err := handler.FSMService.UserProcessPhone(update.Message.Chat.ID, update.Message.Text)
		if err != nil {
			log.Println(err)
			_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
				tu.ID(update.Message.Chat.ID),
				fmt.Sprintf("Произошла ошибка!"),
			))
			return err
		}
	}
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
