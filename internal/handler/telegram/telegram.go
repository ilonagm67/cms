package telegram

import (
	tele "gopkg.in/telebot.v4"
)

var (
	Menu = &tele.ReplyMarkup{ResizeKeyboard: true}
	BtnCatalog = Menu.Text("🛒Каталог")
	BtnOrder = Menu.Text("🛍Заказать товар")
	BtnQuestion = Menu.Text("❓Вопрос")
	BtnNumber = Menu.Text("📞Контакты")
)
