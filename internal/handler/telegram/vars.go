package telegram

import (
	tu "github.com/mymmrac/telego/telegoutil"
)

var MainKeyboard = tu.Keyboard(
	tu.KeyboardRow(
		tu.KeyboardButton("🛒Каталог"),
	),
	tu.KeyboardRow(
		tu.KeyboardButton("🛍Заказать товар"),
	),
	tu.KeyboardRow(
		tu.KeyboardButton("❓Вопрос"),
	),
	tu.KeyboardRow(
		tu.KeyboardButton("📞Контакты"),
	),
)

var NumberKeyboard = tu.Keyboard(
	tu.KeyboardRow(
		tu.KeyboardButton("Отправить номер телефона").WithRequestContact(),
	),
)
