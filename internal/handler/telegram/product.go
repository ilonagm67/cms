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
	tu "github.com/mymmrac/telego/telegoutil"
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
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
			tu.ID(update.Message.Chat.ID),
			fmt.Sprintf("Произошла ошибка"),
		))
		return err
	}
	listjson, err := handler.ProductService.List()
	if err != nil {
		if admin != update.Message.Chat.ID {
			_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
				tu.ID(update.Message.Chat.ID),
				fmt.Sprintf("Произошла ошибка"),
			))
			return err
		}
	} else {
		var list map[string]map[int]*entity.Product
		err = json.Unmarshal(listjson, &list)
		if err != nil {
			_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
				tu.ID(update.Message.Chat.ID),
				fmt.Sprintf("Произошла ошибка"),
			))
			return err
		}
		for Product := range list {
			_, err = ctx.Bot().SendMessage(ctx, tu.Message(
				tu.ID(update.Message.Chat.ID),
				fmt.Sprintf(Product),
			))
			if err != nil {
				return err
			}
		}
	}

	text, err := handler.FSMService.ProductCatalogStart(update.Message.Chat.ID)
	if err != nil {
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
			tu.ID(update.Message.Chat.ID),
			fmt.Sprintf(text),
		))
		return err
	}
	if admin != update.Message.Chat.ID {
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
			tu.ID(update.Message.Chat.ID),
			fmt.Sprintf(text),
		))
	} else {
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
			tu.ID(update.Message.Chat.ID),
			fmt.Sprintf(text),
		).WithReplyMarkup(AdminCatalogKeyboard).WithProtectContent())
	}
	return nil
}

func (handler *ProductHandler) HandleCatalogName(ctx *th.Context, update telego.Update) error {
	adminstr := os.Getenv("ADMIN")
	admin, err := strconv.ParseInt(adminstr, 10, 64)
	if err != nil {
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
			tu.ID(update.Message.Chat.ID),
			fmt.Sprintf("Ошибка"),
		))
		return err
	}
	switch update.Message.Text {
	case "Добавить товар":
		if admin != update.Message.Chat.ID {
			_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
				tu.ID(update.Message.Chat.ID),
				fmt.Sprintf("Ошибка"),
			))
			userjson, err := handler.UserService.Get(update.Message.Chat.ID)
			if err != nil {
				_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
					tu.ID(update.Message.Chat.ID),
					fmt.Sprintf("Ошибка"),
				))
				return err
			}
			var user entity.User
			err = json.Unmarshal(userjson, &user)
			if err != nil {
				_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
					tu.ID(update.Message.Chat.ID),
					fmt.Sprintf("Ошибка"),
				))
				return err
			}
			return errors.New("ID:%v, UserName:%s, Phone: %s, Want to create product")
		} else {
			text, err := handler.FSMService.ProductCatalogName(update.Message.Chat.ID, update.Message.Text)
			if err != nil {
				_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
					tu.ID(update.Message.Chat.ID),
					fmt.Sprintf(text),
				))
				return err
			}
			_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
				tu.ID(update.Message.Chat.ID),
				fmt.Sprintf(text),
			))
			return nil
		}
	default:
		text, err := handler.FSMService.ProductCatalogName(update.Message.Chat.ID, update.Message.Text)
		if err != nil {
			_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
				tu.ID(update.Message.Chat.ID),
				fmt.Sprintf(text),
			))
			return err
		}
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
			tu.ID(update.Message.Chat.ID),
			fmt.Sprintf(text),
		))
		return nil
	}
}

func (handler *ProductHandler) HandleCatalogWeight(ctx *th.Context, update telego.Update) error {
	product, err := handler.FSMService.ProductCatalogWeight(update.Message.Chat.ID, update.Message.Text)
	if err != nil {
		_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
			tu.ID(update.Message.Chat.ID),
			fmt.Sprintf("Произошла ошибка!"),
		))
		return err
	}
	_, _ = ctx.Bot().SendMessage(ctx, tu.Message(
		tu.ID(update.Message.Chat.ID),
		fmt.Sprintf("Название: %s,\nОписание: %s\nВес: %v,\nЦена: %v", product.Name, product.Description, product.Weight, product.Price),
	))
	return nil
}
