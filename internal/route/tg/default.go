package tg

import (
	"cms/internal/handler/telegram"
	"context"
	"log"
	"os"

	"github.com/mymmrac/telego"
)

type Router struct {
	Bot             *telego.Bot
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

	bot, err := telego.NewBot(token, telego.WithDefaultDebugLogger())
	if err != nil {
		log.Fatal(err)
	}

	return &Router{Bot: bot, OrderHandler: Oh, ProductHandler: Ph, UserHandler: Uh, QuestionHandler: Qh}
}

func (r *Router) Init() {
	updates, _ := r.Bot.UpdatesViaLongPolling(context.Background(), nil)

	for update := range updates {
		log.Println(update)
	}
}
