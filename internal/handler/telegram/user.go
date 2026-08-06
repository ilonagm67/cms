package telegram

import (
	"cms/internal/service"
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
	handler.UserService.Add(user.ID,user.FirstName)
	Menu.Reply(
		Menu.Row(BtnCatalog),
		Menu.Row(BtnOrder),
		Menu.Row(BtnQuestion),
		Menu.Row(BtnNumber),
	)
	return c.Send("Добро пожаловать в наш магазин!",Menu)
}
