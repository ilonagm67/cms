package telegram

import (
	"cms/internal/service"
)

type ProductHandler struct {
	FSMService *service.FSMService
}

func NewProductHandler(FSMService *service.FSMService) *ProductHandler {
	return &ProductHandler{FSMService: FSMService}
}
