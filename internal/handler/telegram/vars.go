package telegram

import (
	tele "gopkg.in/telebot.v4"
)

var (
	MainMenu    = &tele.ReplyMarkup{ResizeKeyboard: true}
	BtnCatalog  = MainMenu.Text("🛒Каталог")
	BtnOrder    = MainMenu.Text("🛍Заказать товар")
	BtnQuestion = MainMenu.Text("❓Вопрос")
	BtnNumber   = MainMenu.Text("📞Контакты")
)
