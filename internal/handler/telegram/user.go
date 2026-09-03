package telegram

import (
	"cms/internal/service"
	"fmt"

	tele "gopkg.in/telebot.v4"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(UserService *service.UserService) *UserHandler {
	return &UserHandler{UserService: UserService}
}

func (handler *UserHandler) Start(c tele.Context) error {
	user := c.Sender()
	err := handler.UserService.Add(user.ID, user.FirstName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	MainMenu.Reply(
		MainMenu.Row(BtnCatalog),
		MainMenu.Row(BtnOrder),
		MainMenu.Row(BtnQuestion),
		MainMenu.Row(BtnNumber),
	)
	return c.Send("Добро пожаловать в наш магазин!", MainMenu)
}

func (handler *UserHandler) Contact(c tele.Context) error {
	return nil
}
