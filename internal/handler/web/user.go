package web

import (
	"cms/internal/service"
	"encoding/json"
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
		_, _ = fmt.Fprintf(w, "Error")
	} else {
		user, err := handler.UserService.Get(id)
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

func (handler *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	list, err := handler.UserService.List()
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

func (handler *UserHandler) Health(w http.ResponseWriter, r *http.Request) {
	message, err := handler.UserService.Health()
	if err != nil {
		logging := fmt.Sprintf("IP: %s, Path: %s, Error: %s", r.RemoteAddr, r.RequestURI, err.Error())
		log.Println(logging)
		_, _ = fmt.Fprintf(w, "Error")
	} else {
		json, err := json.Marshal(message)
		if err != nil {
			log.Println(err)
			_, _ = fmt.Fprintf(w, "Error")
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintf(w, "%s", json)
	}
}
