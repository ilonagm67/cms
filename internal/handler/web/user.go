package web

import (
	"cms/internal/service"
	"net/http"
	"strconv"
	"fmt"
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
		fmt.Fprintf(w,err.Error())
	}
	user,err := handler.UserService.Get(id)
	if err != nil {
		fmt.Fprintf(w,err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w,string(user))
}

func(handler *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	list,err := handler.UserService.List()
	if err != nil {
		fmt.Fprintf(w,err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w,string(list))
}
