package tg

import (
	"os"
	"time"
	"log"
	"cms/internal/handler/telegram"
	tele "gopkg.in/telebot.v4"
)

type Router struct {
	Oh *telegram.OrderHandler
	Ph *telegram.ProductHandler
	Uh *telegram.UserHandler
}

func NewRouter(Oh *telegram.OrderHandler, Ph *telegram.ProductHandler, Uh *telegram.UserHandler) {
	pref := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start",Uh.Start)

	b.Start()
}
