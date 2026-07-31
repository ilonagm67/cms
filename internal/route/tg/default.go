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

var (
	menu = &tele.ReplyMarkup{ResizeKeyboard: true}
	btnCatalog = menu.Text("🛒Каталог")
	btnQuestion = menu.Text("❓Вопрос")
	btnNumber = menu.Text("📞Контакты")
)

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

	menu.Reply(
		menu.Row(btnCatalog),
		menu.Row(btnOrder),
		menu.Row(btnQuestion),
		menu.Row(btnNumber),
	)

	b.Handle("/start",func (c tele.Context) error {
		return c.Send("Добро пожаловать в наш магазин!",menu)
	})

	b.Start()
}
