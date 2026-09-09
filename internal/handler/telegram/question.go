package telegram

import (
	"cms/internal/service"
)

type QuestionHandler struct {
	UserService *service.UserService
}

func NewQuestionHandler(UserService *service.UserService) *QuestionHandler {
	return &QuestionHandler{UserService: UserService}
}
