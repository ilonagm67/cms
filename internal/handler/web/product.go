package web

import (
	"cms/internal/service"
	"fmt"
	"log"
	"net/http"
)

type ProductHandler struct {
	ProductService *service.ProductService
}

func NewProductHandler(ProductService *service.ProductService) *ProductHandler {
	return &ProductHandler{ProductService: ProductService}
}

func (handler *ProductHandler) List(w http.ResponseWriter, r *http.Request) {
	handler.ProductService.Add("test", 50)
	list, err := handler.ProductService.List()
	if err != nil {
		logging := fmt.Sprintf("IP: %s, Path: %s, Error: %s", r.RemoteAddr, r.RequestURI, err.Error())
		log.Println(logging)
		fmt.Fprintf(w, "Error")
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", list)
}
