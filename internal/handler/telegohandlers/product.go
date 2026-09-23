package telegohandlers

import (
	"cms/internal/service"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

type ProductHandler struct {
	FSMService     *service.FSMService
	ProductService *service.ProductService
}

func NewProductHandler(FSMService *service.FSMService, ProductService *service.ProductService) *ProductHandler {
	return &ProductHandler{FSMService: FSMService, ProductService: ProductService}
}

func (handler *ProductHandler) HandleStart(ctx *th.Context, update telego.Update) error {
	admin, err := handler.FSMService.ProductCatalogStart(update.Message.Chat.ID)
	if err != nil {
		SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", MainKeyboard)
		return err
	}
	var rows [][]telego.KeyboardButton
	var row []telego.KeyboardButton
	count := 0
	list, _ := handler.ProductService.List()
	for Product := range list {
		row = append(row, tu.KeyboardButton(Product))
		count++
		if count%2 == 0 {
			rows = append(rows, row)
			row = nil
		}
	}
	if len(row) > 0 {
		rows = append(rows, row)
	}
	if !admin {
		keyboard := tu.Keyboard(rows...).WithResizeKeyboard().WithOneTimeKeyboard()
		SendMessage(ctx, update.Message.Chat.ID, "Выберите опцию:", keyboard)
		return err
	} else {
		row = tu.KeyboardRow(tu.KeyboardButton("Добавить товар"))
		rows = append(rows, row)
		keyboard := tu.Keyboard(rows...).WithResizeKeyboard().WithOneTimeKeyboard()
		SendMessage(ctx, update.Message.Chat.ID, "Выберите опцию:", keyboard)
		return nil
	}
}

func (handler *ProductHandler) HandleCatalogName(ctx *th.Context, update telego.Update) error {
	adminstr := os.Getenv("ADMIN")
	admin, err := strconv.ParseInt(adminstr, 10, 64)
	if err != nil {
		return err
	}
	switch update.Message.Text {
	case "Добавить товар":
		if admin != update.Message.Chat.ID {
			SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", MainKeyboard)
			return errors.New("Want to add product!")
		} else {
			err = handler.FSMService.ProductCatalogName(update.Message.Chat.ID, update.Message.Text)
			SendMessage(ctx, update.Message.Chat.ID, "Введите название:", nil)
			return nil
		}
	default:
		err := handler.FSMService.ProductCatalogName(update.Message.Chat.ID, update.Message.Text)
		if err != nil {
			SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", MainKeyboard)
			return err
		}
		var rows [][]telego.KeyboardButton
		var row []telego.KeyboardButton
		count := 0
		list, _ := handler.ProductService.List()
		for _, Weights := range list {
			for Weight, Product := range Weights {
				if Product.Name == update.Message.Text {
					weight := strconv.Itoa(Weight)
					text := fmt.Sprintf("%dкг", Product.Weight)
					err := SendMessage(ctx, update.Message.Chat.ID, text, nil)
					if err != nil {
						return err
					}
					row = append(row, tu.KeyboardButton(weight))
					count++
					if count%2 == 0 {
						rows = append(rows, row)
						row = nil
					}
				}
			}
		}
		if len(row) > 0 {
			rows = append(rows, row)
		}
		keyboard := tu.Keyboard(rows...).WithResizeKeyboard().WithOneTimeKeyboard()
		SendMessage(ctx, update.Message.Chat.ID, "Выберите вес:", keyboard)
		return nil
	}
}

func (handler *ProductHandler) HandleCatalogWeight(ctx *th.Context, update telego.Update) error {
	weight, err := strconv.Atoi(update.Message.Text)
	if err != nil {
		SendMessage(ctx, update.Message.Chat.ID, "Введите нормальное число!", nil)
		return err
	}
	product, err := handler.FSMService.ProductCatalogWeight(update.Message.Chat.ID, weight)
	if err != nil {
		SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", MainKeyboard)
		return err
	}
	text := fmt.Sprintf("Название: %s\n\nОписание: %s\n\nВес: %d\nЦена: %d", product.Name, product.Description, product.Weight, product.Price)
	SendMessagePhoto(ctx, update.Message.Chat.ID, text, product.Image, MainKeyboard)
	return nil
}

func (handler *ProductHandler) HandleCatalogAddName(ctx *th.Context, update telego.Update) error {
	err := handler.FSMService.ProductAddingName(update.Message.Chat.ID, update.Message.Text)
	if err != nil {
		SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", MainKeyboard)
		return err
	}
	SendMessage(ctx, update.Message.Chat.ID, "Введите вес:", nil)
	return nil
}

func (handler *ProductHandler) HandleCatalogAddWeight(ctx *th.Context, update telego.Update) error {
	weight, err := strconv.Atoi(update.Message.Text)
	if err != nil {
		SendMessage(ctx, update.Message.Chat.ID, "Введите нормальное число!", nil)
		return err
	}
	err = handler.FSMService.ProductAddingWeight(update.Message.Chat.ID, weight)
	if err != nil {
		SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", MainKeyboard)
		return err
	}
	SendMessage(ctx, update.Message.Chat.ID, "Введите Описание:", nil)
	return nil
}

func (handler *ProductHandler) HandleCatalogAddDescription(ctx *th.Context, update telego.Update) error {
	err := handler.FSMService.ProductAddingDescription(update.Message.Chat.ID, update.Message.Text)
	if err != nil {
		SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", MainKeyboard)
		return err
	}
	SendMessage(ctx, update.Message.Chat.ID, "Отправьте изображение:", nil)
	return nil
}

func (handler *ProductHandler) HandleCatalogAddImage(ctx *th.Context, update telego.Update) error {
	switch {
	case len(update.Message.Photo) > 0:
		chatID := update.Message.Chat.ID
		photos := update.Message.Photo

		if len(photos) > 0 {
			bestPhoto := photos[len(photos)-1]

			file, err := DownloadImage(ctx, "photo", bestPhoto.FileID, bestPhoto.FileUniqueID)
			if err != nil {
				SendMessage(ctx, chatID, "Произошла ошибка", MainKeyboard)
				return err
			}

			err = handler.FSMService.ProductAddingImage(chatID, file)
			if err != nil {
				SendMessage(ctx, chatID, "Произошла ошибка", MainKeyboard)
				return err
			}

			SendMessage(ctx, chatID, "Введите цену:", nil)
			return nil
		}
	default:
		SendMessage(ctx, update.Message.Chat.ID, "Отправьте Изображение!", nil)
	}
	return nil
}

func (handler *ProductHandler) HandleCatalogAddPrice(ctx *th.Context, update telego.Update) error {
	price, err := strconv.Atoi(update.Message.Text)
	if err != nil {
		SendMessage(ctx, update.Message.Chat.ID, "Введите нормальное число!", nil)
		return err
	}
	err = handler.FSMService.ProductAddingPrice(update.Message.Chat.ID, price)
	if err != nil {
		SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", MainKeyboard)
		return err
	}
	SendMessage(ctx, update.Message.Chat.ID, "Товар создан!", MainKeyboard)
	return nil
}
