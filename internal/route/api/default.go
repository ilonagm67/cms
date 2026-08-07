package api

import (
	"cms/internal/handler/web"
	"net/http"
)

type Router struct {
	Oh *web.OrderHandler
	Ph *web.ProductHandler
	Uh *web.UserHandler
}

func NewRouter(Oh *web.OrderHandler, Ph *web.ProductHandler, Uh *web.UserHandler) {
	http.HandleFunc("/user/{id}",Uh.Get)
	http.HandleFunc("/users",Uh.List)
	http.ListenAndServe(":8080",nil)
}
