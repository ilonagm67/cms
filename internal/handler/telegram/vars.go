package telegram

import (
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
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
).WithOneTimeKeyboard().WithResizeKeyboard()

var NumberKeyboard = tu.Keyboard(
	tu.KeyboardRow(
		tu.KeyboardButton("📞Отправить номер телефона").WithRequestContact(),
	),
).WithOneTimeKeyboard().WithResizeKeyboard()

var AdminCatalogKeyboard = tu.Keyboard(
	tu.KeyboardRow(
		tu.KeyboardButton("Добавить товар"),
	),
).WithOneTimeKeyboard().WithResizeKeyboard()

var DeliveryKeyboard = tu.Keyboard(
	tu.KeyboardRow(
		tu.KeyboardButton("Новая Почта"),
		tu.KeyboardButton("Укр Почта"),
	),
	tu.KeyboardRow(
		tu.KeyboardButton("Самовывоз"),
		tu.KeyboardButton("Доставка"),
	),
).WithOneTimeKeyboard().WithResizeKeyboard()

var DeliveryBaseKeyboard = tu.Keyboard(
	tu.KeyboardRow(
		tu.KeyboardButton("База"),
	),
).WithOneTimeKeyboard().WithResizeKeyboard()

var PayTypeKeyboard = tu.Keyboard(
	tu.KeyboardRow(
		tu.KeyboardButton("Перевод на карту"),
		tu.KeyboardButton("Оплата при получении"),
	),
).WithOneTimeKeyboard().WithResizeKeyboard()

var ReturnKeyboard = tu.Keyboard(
	tu.KeyboardRow(
		tu.KeyboardButton("🔙Назад"),
	),
).WithOneTimeKeyboard()

func SendMessage(ctx *th.Context, ID int64, text string, keyboard telego.ReplyMarkup) error {
	if keyboard != nil {
		_, err := ctx.Bot().SendMessage(ctx, tu.Message(
			tu.ID(ID),
			text,
		).WithReplyMarkup(keyboard))
		return err
	} else {
		_, err := ctx.Bot().SendMessage(ctx, tu.Message(
			tu.ID(ID),
			text,
		))
		return err
	}
}
