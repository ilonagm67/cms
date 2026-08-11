package api

import (
	"cms/internal/handler/web"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	Router         *mux.Router
	OrderHandler   *web.OrderHandler
	ProductHandler *web.ProductHandler
	UserHandler    *web.UserHandler
}

func NewRouter(Oh *web.OrderHandler, Ph *web.ProductHandler, Uh *web.UserHandler) *Router {
	r := mux.NewRouter()
	return &Router{Router: r, OrderHandler: Oh, ProductHandler: Ph, UserHandler: Uh}
}

func (r *Router) Init() {
	r.Router.HandleFunc("/api/user/{id}", r.UserHandler.Get)
	r.Router.HandleFunc("/api/users", r.UserHandler.List)

	r.Router.HandleFunc("/api/products", r.ProductHandler.List)

	r.Router.HandleFunc("/api/order/{id}", r.OrderHandler.Get)
	r.Router.HandleFunc("/api/orders", r.OrderHandler.List)

	http.Handle("/", r.Router)
	http.ListenAndServe(":8080", nil)
}
