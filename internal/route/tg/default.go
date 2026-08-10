package tg

import (
	"cms/internal/handler/telegram"
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v4"
)

type Router struct {
	Oh *telegram.OrderHandler
	Ph *telegram.ProductHandler
	Uh *telegram.UserHandler
	Qh *telegram.QuestionHandler
}

func NewRouter(Oh *telegram.OrderHandler, Ph *telegram.ProductHandler, Uh *telegram.UserHandler, Qh *telegram.QuestionHandler) {
	pref := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", Uh.Start)
	b.Handle("🛒Каталог", Ph.List)

	b.Start()
}
