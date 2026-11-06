package telegohandlers

import (
	"os"

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
		tu.KeyboardButton("📦Мой заказ"),
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

var PreDeliveryKeyboard = tu.Keyboard(
	tu.KeyboardRow(
		tu.KeyboardButton("Добавить товар"),
	),
	tu.KeyboardRow(
		tu.KeyboardButton("Удалить товар"),
	),
	tu.KeyboardRow(
		tu.KeyboardButton("Оформить доставку"),
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
).WithOneTimeKeyboard().WithResizeKeyboard()

func SendMessagePhoto(ctx *th.Context, ID int64, text string, photo string, keyboard telego.ReplyMarkup) error {
	_, err := ctx.Bot().SendPhoto(ctx, tu.Photo(
		tu.ID(ID),
		tu.File(OpenPhoto(photo)),
	).WithCaption(text).WithReplyMarkup(keyboard))
	return err
}

func OpenPhoto(photo string) *os.File {
	file, _ := os.Open(photo)
	return file
}

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
