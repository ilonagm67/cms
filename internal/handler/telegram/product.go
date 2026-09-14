package telegram

import (
	"cms/internal/entity"
	"cms/internal/service"
	"encoding/json"
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
		SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", nil)
		return err
	}
	listjson, err := handler.ProductService.List()
	if err != nil {
		if admin != update.Message.Chat.ID {
			SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", nil)
			return err
		}
	} else {
		var list map[string]map[int]*entity.Product
		err = json.Unmarshal(listjson, &list)
		if err != nil {
			SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", nil)
			return err
		}
		for Product := range list {
			err := SendMessage(ctx, update.Message.Chat.ID, Product, nil)
			if err != nil {
				return err
			}
		}
	}

	text, err := handler.FSMService.ProductCatalogStart(update.Message.Chat.ID)
	if err != nil {
		SendMessage(ctx, update.Message.Chat.ID, text, nil)
		return err
	}
	if admin != update.Message.Chat.ID {
		SendMessage(ctx, update.Message.Chat.ID, text, nil)
	} else {
		SendMessage(ctx, update.Message.Chat.ID, text, AdminCatalogKeyboard)
	}
	return nil
}

func (handler *ProductHandler) HandleCatalogName(ctx *th.Context, update telego.Update) error {
	adminstr := os.Getenv("ADMIN")
	admin, err := strconv.ParseInt(adminstr, 10, 64)
	if err != nil {
		SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", nil)
		return err
	}
	switch update.Message.Text {
	case "Добавить товар":
		if admin != update.Message.Chat.ID {
			SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", nil)
			userjson, err := handler.UserService.Get(update.Message.Chat.ID)
			if err != nil {
				SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", nil)
				return err
			}
			var user entity.User
			err = json.Unmarshal(userjson, &user)
			if err != nil {
				SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", nil)
				return err
			}
			//err = fmt.Sprintf("ID: %d, UserName: %s, Phone: %s, Want to create product", user.ID, user.Name, user.Number)
			return errors.New("Want to Create Product")
		}
		text, err := handler.FSMService.ProductCatalogName(update.Message.Chat.ID, update.Message.Text)
		if err != nil {
			SendMessage(ctx, update.Message.Chat.ID, text, nil)
			return err
		}
		SendMessage(ctx, update.Message.Chat.ID, text, nil)
		return nil
	default:
		text, err := handler.FSMService.ProductCatalogName(update.Message.Chat.ID, update.Message.Text)
		if err != nil {
			SendMessage(ctx, update.Message.Chat.ID, text, nil)
			return err
		}
		SendMessage(ctx, update.Message.Chat.ID, text, nil)
		encoded, err := handler.ProductService.List()
		if err != nil {
			SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", nil)
			return err
		}
		var list map[string]map[int]*entity.Product
		err = json.Unmarshal(encoded, &list)
		if err != nil {
			SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", nil)
			return err
		}
		for ProductName, Weights := range list {
			for Weight, _ := range Weights {
				if ProductName == update.Message.Text {
					text = fmt.Sprintf("%d", Weight)
					err := SendMessage(ctx, update.Message.Chat.ID, text, nil)
					if err != nil {
						return err
					}
				}
			}
		}
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
		SendMessage(ctx, update.Message.Chat.ID, "Произошла ошибка", nil)
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
