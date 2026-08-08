package web

import (
	"cms/internal/service"
	"net/http"
	"strconv"
	"fmt"
	"log"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(UserService *service.UserService) *UserHandler {
	return &UserHandler{UserService: UserService}
}

func(handler *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	id,err := strconv.ParseInt(r.PathValue("id"),10,64)
	if err != nil {
		logging := fmt.Sprintf("IP: %s, Path: %s, Error: %s",r.RemoteAddr,r.RequestURI,err.Error())
		log.Println(logging)
		fmt.Fprintf(w,"Error")
	}
	user,err := handler.UserService.Get(id)
	if err != nil {
		logging := fmt.Sprintf("IP: %s, Path: %s, Error: %s",r.RemoteAddr,r.RequestURI,err.Error())
		log.Println(logging)
		fmt.Fprintf(w,"Error")
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w,string(user))
}

func(handler *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	list,err := handler.UserService.List()
	if err != nil {
		logging := fmt.Sprintf("IP: %s, Path: %s, Error: %s",r.RemoteAddr,r.RequestURI,err.Error())
		log.Println(logging)
		fmt.Fprintf(w,"Error")
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w,string(list))
}
