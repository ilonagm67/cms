package tg

import (
	"cms/internal/handler/telegram"
	"log"
	"os"
)

type Router struct {
	OrderHandler    *telegram.OrderHandler
	ProductHandler  *telegram.ProductHandler
	UserHandler     *telegram.UserHandler
	QuestionHandler *telegram.QuestionHandler
}

func NewRouter(Oh *telegram.OrderHandler, Ph *telegram.ProductHandler, Uh *telegram.UserHandler, Qh *telegram.QuestionHandler) *Router {
	token := os.Getenv("TOKEN")
	if token == "" {
		log.Fatal("env TOKEN not found")
	}

	return &Router{OrderHandler: Oh, ProductHandler: Ph, UserHandler: Uh, QuestionHandler: Qh}
}

func (r *Router) Init() {}
