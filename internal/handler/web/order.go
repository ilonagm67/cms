package web

import (
	"cms/internal/service"
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
		_, err = fmt.Fprintf(w, "Error")
		if err != nil {
			log.Println(err)
		}
	} else {
		user, err := handler.OrderService.Get(id)
		if err != nil {
			logging := fmt.Sprintf("IP: %s, Path: %s, Error: %s", r.RemoteAddr, r.RequestURI, err.Error())
			log.Println(logging)
			_, err = fmt.Fprintf(w, "Error")
			if err != nil {
				log.Println(err)
			}
		} else {
			w.Header().Set("Content-Type", "application/json")
			_, err := fmt.Fprintf(w, "%s", user)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func (handler *OrderHandler) List(w http.ResponseWriter, r *http.Request) {
	list, err := handler.OrderService.List()
	if err != nil {
		logging := fmt.Sprintf("IP: %s, Path: %s, Error: %s", r.RemoteAddr, r.RequestURI, err.Error())
		log.Println(logging)
		_, err = fmt.Fprintf(w, "Error")
		if err != nil {
			log.Println(err)
		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		_, err = fmt.Fprintf(w, "%s", list)
		if err != nil {
			log.Println(err)
		}
	}
}
