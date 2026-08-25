package telegram

import (
	"cms/internal/service"

	tele "gopkg.in/telebot.v4"
)

type QuestionHandler struct {
	UserService *service.UserService
}

func NewQuestionHandler(UserService *service.UserService) *QuestionHandler {
	return &QuestionHandler{UserService: UserService}
}

func (handler *QuestionHandler) Start(c tele.Context) error {
	return nil
}
