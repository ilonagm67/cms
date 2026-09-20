package telegohandlers

import (
	"io"
	"net/http"
	"os"
	"path"

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
).WithOneTimeKeyboard()

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

func DownloadImage(ctx *th.Context, Folder, ID, Unique string) (string, error) {
	file, err := ctx.Bot().GetFile(ctx, &telego.GetFileParams{
		FileID: ID,
	})
	if err != nil {
		return "", err
	}
	URL := ctx.Bot().FileDownloadURL(file.FilePath)
	resp, err := http.Get(URL)
	if err != nil {
		return "", err
	}
	_, err = os.ReadDir(Folder)
	if err != nil {
		err = os.Mkdir(Folder, 0755)
		if err != nil {
			return "", err
		}
	}
	filepath := path.Join(Folder, Unique+".jpg")
	out, err := os.Create(filepath)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}
	return filepath, nil
}
