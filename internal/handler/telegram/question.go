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

func (handler *QuestionHandler) Middleware(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		err := next(c)
		return err
	}
}

func (handler *QuestionHandler) Start(c tele.Context) error {
	return nil
}
