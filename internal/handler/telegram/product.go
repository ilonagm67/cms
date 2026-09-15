package telegram

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
	UserService    *service.UserService
}

func NewProductHandler(FSMService *service.FSMService, ProductService *service.ProductService, UserService *service.UserService) *ProductHandler {
	return &ProductHandler{FSMService: FSMService, ProductService: ProductService, UserService: UserService}
}

func (handler *ProductHandler) HandleStart(ctx *th.Context, update telego.Update) error {
	adminstr := os.Getenv("ADMIN")
	admin, err := strconv.ParseInt(adminstr, 10, 64)
	if err != nil {
		return err
	}
	err = handler.FSMService.ProductCatalogStart(update.Message.Chat.ID)
	if err != nil {
		SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", MainKeyboard)
		return err
	}
	list, err := handler.ProductService.List()
	if err != nil {
		if admin != update.Message.Chat.ID {
			SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", MainKeyboard)
			return err
		} else {
			SendMessage(ctx, update.Message.Chat.ID, "Выберите опцию:", AdminCatalogKeyboard)
			return nil
		}
	}
	for Product := range list {
		err = SendMessage(ctx, update.Message.Chat.ID, Product, nil)
		if err != nil {
			return err
		}
	}
	if admin != update.Message.Chat.ID {
		SendMessage(ctx, update.Message.Chat.ID, "Выберите опцию:", nil)
		return err
	} else {
		SendMessage(ctx, update.Message.Chat.ID, "Выберите опцию:", AdminCatalogKeyboard)
		return nil
	}
}

func (handler *ProductHandler) HandleCatalogName(ctx *th.Context, update telego.Update) error {
	adminstr := os.Getenv("ADMIN")
	admin, err := strconv.ParseInt(adminstr, 10, 64)
	if err != nil {
		return err
	}
	if update.Message.Text == "Добавить товар" {
		if admin != update.Message.Chat.ID {
			SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", MainKeyboard)
			return errors.New("Want to add product!")
		}
	}
	catalog, err := handler.ProductService.List()
	if admin != update.Message.Chat.ID {
		if err != nil {
			SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", MainKeyboard)
			return err
		}
		for product := range catalog {
			if product != update.Message.Text {
				SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", MainKeyboard)
				return errors.New("Product Not Found")
			}
		}
	}
	err = handler.FSMService.ProductCatalogName(update.Message.Chat.ID, update.Message.Text)
	if err != nil {
		SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", MainKeyboard)
		return err
	}
	if admin != update.Message.Chat.ID {
		SendMessage(ctx, update.Message.Chat.ID, "Выберите вес:", MainKeyboard)
	} else {
		SendMessage(ctx, update.Message.Chat.ID, "Введите название:", nil)
	}
	return nil
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
	text, err := handler.FSMService.ProductAddingName(update.Message.Chat.ID, update.Message.Text)
	if err != nil {
		SendMessage(ctx, update.Message.Chat.ID, text, nil)
		return err
	}
	SendMessage(ctx, update.Message.Chat.ID, text, nil)
	return nil
}

func (handler *ProductHandler) HandleCatalogAddWeight(ctx *th.Context, update telego.Update) error {
	weight, err := strconv.Atoi(update.Message.Text)
	if err != nil {
		SendMessage(ctx, update.Message.Chat.ID, "Введите нормальное число!", nil)
		return err
	}
	text, err := handler.FSMService.ProductAddingWeight(update.Message.Chat.ID, weight)
	if err != nil {
		SendMessage(ctx, update.Message.Chat.ID, text, nil)
		return err
	}
	SendMessage(ctx, update.Message.Chat.ID, text, nil)
	return nil
}

func (handler *ProductHandler) HandleCatalogAddDescription(ctx *th.Context, update telego.Update) error {
	text, err := handler.FSMService.ProductAddingDescription(update.Message.Chat.ID, update.Message.Text)
	if err != nil {
		SendMessage(ctx, update.Message.Chat.ID, text, nil)
		return err
	}
	SendMessage(ctx, update.Message.Chat.ID, text, nil)
	return nil
}

func (handler *ProductHandler) HandleCatalogAddImage(ctx *th.Context, update telego.Update) error {
	text, err := handler.FSMService.ProductAddingImage(update.Message.Chat.ID, update.Message.Text)
	if err != nil {
		SendMessage(ctx, update.Message.Chat.ID, text, nil)
		return err
	}
	SendMessage(ctx, update.Message.Chat.ID, text, nil)
	return nil
}

func (handler *ProductHandler) HandleCatalogAddPrice(ctx *th.Context, update telego.Update) error {
	price, err := strconv.Atoi(update.Message.Text)
	if err != nil {
		SendMessage(ctx, update.Message.Chat.ID, "Введите нормальное число!", nil)
		return err
	}
	text, err := handler.FSMService.ProductAddingPrice(update.Message.Chat.ID, price)
	if err != nil {
		SendMessage(ctx, update.Message.Chat.ID, text, nil)
		return err
	}
	SendMessage(ctx, update.Message.Chat.ID, text, MainKeyboard)
	return nil
}
