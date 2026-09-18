package telegohandlers

import (
	"cms/internal/service"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

type ProductHandler struct {
	FSMService     *service.FSMService
	ProductService *service.ProductService
}

func NewProductHandler(FSMService *service.FSMService, ProductService *service.ProductService) *ProductHandler {
	return &ProductHandler{FSMService: FSMService, ProductService: ProductService}
}

func (handler *ProductHandler) HandleStart(ctx *th.Context, update telego.Update) error {
	adminstr := os.Getenv("ADMIN")
	admin, err := strconv.ParseInt(adminstr, 10, 64)
	if err != nil {
		return err
	}
	list, err := handler.ProductService.List()
	if err != nil {
		if admin != update.Message.Chat.ID {
			err = handler.FSMService.ProductCatalogStart(update.Message.Chat.ID, false)
			SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", MainKeyboard)
			return err
		} else {
			err = handler.FSMService.ProductCatalogStart(update.Message.Chat.ID, true)
			SendMessage(ctx, update.Message.Chat.ID, "Выберите опцию:", AdminCatalogKeyboard)
			return nil
		}
	} else {
		for Product := range list {
			err = SendMessage(ctx, update.Message.Chat.ID, Product, nil)
			if err != nil {
				return err
			}
		}
		if admin != update.Message.Chat.ID {
			err = handler.FSMService.ProductCatalogStart(update.Message.Chat.ID, false)
			SendMessage(ctx, update.Message.Chat.ID, "Выберите опцию:", nil)
			return err
		} else {
			err = handler.FSMService.ProductCatalogStart(update.Message.Chat.ID, true)
			SendMessage(ctx, update.Message.Chat.ID, "Выберите опцию:", AdminCatalogKeyboard)
			return nil
		}
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
		list, _ := handler.ProductService.List()
		for _, Weights := range list {
			for Weight, Product := range Weights {
				if Product.Name == update.Message.Text {
					weight := strconv.Itoa(Weight)
					err := SendMessage(ctx, update.Message.Chat.ID, weight, nil)
					if err != nil {
						return err
					}
				}
			}
		}
		SendMessage(ctx, update.Message.Chat.ID, "Выберите вес:", nil)
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
	SendMessage(ctx, update.Message.Chat.ID, text, MainKeyboard)
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
	SendMessage(ctx, update.Message.Chat.ID, "Назовите Изображение:", nil)
	return nil
}

func (handler *ProductHandler) HandleCatalogAddImage(ctx *th.Context, update telego.Update) error {
	err := handler.FSMService.ProductAddingImage(update.Message.Chat.ID, update.Message.Text)
	if err != nil {
		SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", MainKeyboard)
		return err
	}
	SendMessage(ctx, update.Message.Chat.ID, "Введите цену:", nil)
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
