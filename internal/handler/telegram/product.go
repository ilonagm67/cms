package telegram

import (
	"cms/internal/entity"
	"cms/internal/service"
	"encoding/json"
	"fmt"
	"log"

	tele "gopkg.in/telebot.v4"
)

type ProductHandler struct {
	ProductService *service.ProductService
}

func NewProductHandler(ProductService *service.ProductService) *ProductHandler {
	return &ProductHandler{ProductService: ProductService}
}

func (handler *ProductHandler) Middleware(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		err := next(c)
		return err
	}
}

func (handler *ProductHandler) List(c tele.Context) error {
	user := c.Sender()
	list, err := handler.ProductService.List()

	if err != nil {
		logging := fmt.Sprintf("ID: %v, Error: %s", user.ID, err.Error())
		log.Println(logging)
		return c.Send("Произошла ошибка")
	}

	var p map[string]map[int]*entity.Product
	err = json.Unmarshal(list, &p)
	if err != nil {
		logging := fmt.Sprintf("ID: %v, Error: %s", user.ID, err.Error())
		log.Println(logging)
		return c.Send("Произошла ошибка")
	}

	for key, _ := range p {
		err := c.Send(key)
		if err != nil {
			return err
		}
	}
	return nil
}
