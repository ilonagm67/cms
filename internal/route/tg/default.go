package tg

import (
	"cms/internal/handler/telegram"
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v4"
)

type Router struct {
	Bot             *tele.Bot
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

	pref := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &Router{Bot: b, OrderHandler: Oh, ProductHandler: Ph, UserHandler: Uh, QuestionHandler: Qh}
}

func (r *Router) Init() {
	r.Bot.Handle("/start", r.UserHandler.Start)
	r.Bot.Handle("🛒Каталог", r.ProductHandler.List)
	r.Bot.Handle("🛍Заказать товар", r.OrderHandler.Start)
	r.Bot.Handle("❓Вопрос", r.QuestionHandler.Start)
	r.Bot.Handle("📞Контакты", r.UserHandler.Contact)

	r.Bot.Start()
}
