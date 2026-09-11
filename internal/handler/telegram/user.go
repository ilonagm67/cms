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
	UserService *service.UserService
	FSMService  *service.FSMService
}

func NewUserHandler(UserService *service.UserService, FSMService *service.FSMService) *UserHandler {
	return &UserHandler{UserService: UserService, FSMService: FSMService}
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
	handler.UserService.Start(update.Message.Chat.ID)
	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
		tu.ID(update.Message.Chat.ID),
		fmt.Sprintf("Enter Your name: "),
	))
	return nil
}

func (handler *UserHandler) HandleName(ctx *th.Context, update telego.Update) error {
	handler.UserService.ProcessName(update.Message.Chat.ID, update.Message.Text)
	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
		tu.ID(update.Message.Chat.ID),
		fmt.Sprintf("Hello %s!", update.Message.Text),
	))
	return nil
}
