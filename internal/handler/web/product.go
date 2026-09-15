package web

import (
	"cms/internal/service"
	"encoding/json"
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
	list, err := handler.ProductService.List()
	if err != nil {
		logging := fmt.Sprintf("IP: %s, Path: %s, Error: %s", r.RemoteAddr, r.RequestURI, err.Error())
		log.Println(logging)
		_, _ = fmt.Fprintf(w, "Error")
	} else {
		json, err := json.Marshal(list)
		if err != nil {
			log.Println(err)
			_, _ = fmt.Fprintf(w, "Error")
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintf(w, "%s", json)
	}
}
