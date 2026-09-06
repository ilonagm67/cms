package web

import (
	"cms/internal/service"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(UserService *service.UserService) *UserHandler {
	return &UserHandler{UserService: UserService}
}

func (handler *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		logging := fmt.Sprintf("IP: %s, Path: %s, Error: %s", r.RemoteAddr, r.RequestURI, err.Error())
		log.Println(logging)
		_, err = fmt.Fprintf(w, "Error")
		if err != nil {
			log.Println(err)
		}
	}
	user, err := handler.UserService.Get(id)
	if err != nil {
		logging := fmt.Sprintf("IP: %s, Path: %s, Error: %s", r.RemoteAddr, r.RequestURI, err.Error())
		log.Println(logging)
		_, err = fmt.Fprintf(w, "Error")
		if err != nil {
			log.Println(err)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = fmt.Fprintf(w, "%s", user)
	if err != nil {
		log.Println(err)
	}
}

func (handler *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	list, err := handler.UserService.List()
	if err != nil {
		logging := fmt.Sprintf("IP: %s, Path: %s, Error: %s", r.RemoteAddr, r.RequestURI, err.Error())
		log.Println(logging)
		_, err = fmt.Fprintf(w, "Error")
		if err != nil {
			log.Println(err)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = fmt.Fprintf(w, "%s", list)
	if err != nil {
		log.Println(err)
	}
}

func (handler *UserHandler) Health(w http.ResponseWriter, r *http.Request) {
	message, err := handler.UserService.Health()
	if err != nil {
		logging := fmt.Sprintf("IP: %s, Path: %s, Error: %s", r.RemoteAddr, r.RequestURI, err.Error())
		log.Println(logging)
		_, err = fmt.Fprintf(w, "Error")
		if err != nil {
			log.Println(err)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = fmt.Fprintf(w, "%s", message)
	if err != nil {
		log.Println(err)
	}
}
