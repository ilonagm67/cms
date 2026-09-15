package web

import (
	"cms/internal/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type OrderHandler struct {
	OrderService *service.OrderService
}

func NewOrderHandler(OrderService *service.OrderService) *OrderHandler {
	return &OrderHandler{OrderService: OrderService}
}

func (handler *OrderHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		logging := fmt.Sprintf("IP: %s, Path: %s, Error: %s", r.RemoteAddr, r.RequestURI, err.Error())
		log.Println(logging)
		_, _ = fmt.Fprintf(w, "Error")
	} else {
		user, err := handler.OrderService.Get(id)
		if err != nil {
			logging := fmt.Sprintf("IP: %s, Path: %s, Error: %s", r.RemoteAddr, r.RequestURI, err.Error())
			log.Println(logging)
			_, _ = fmt.Fprintf(w, "Error")
		} else {
			json, err := json.Marshal(user)
			if err != nil {
				log.Println(err)
				_, _ = fmt.Fprintf(w, "Error")
			}
			w.Header().Set("Content-Type", "application/json")
			_, _ = fmt.Fprintf(w, "%s", json)
		}
	}
}

func (handler *OrderHandler) List(w http.ResponseWriter, r *http.Request) {
	list, err := handler.OrderService.List()
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
