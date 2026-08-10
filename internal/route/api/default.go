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
	http.HandleFunc("/api/user/{id}", Uh.Get)
	http.HandleFunc("/api/users", Uh.List)

	http.HandleFunc("/api/products", Ph.List)

	http.HandleFunc("/api/order/{id}", Oh.Get)
	http.HandleFunc("/api/orders", Oh.List)

	http.ListenAndServe(":8080", nil)
}
