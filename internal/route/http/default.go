package http

import (
	"cms/internal/handler/web"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Router struct {
	Router         *mux.Router
	Metrics        *prometheus.Registry
	OrderHandler   *web.OrderHandler
	ProductHandler *web.ProductHandler
	UserHandler    *web.UserHandler
}

func NewRouter(Oh *web.OrderHandler, Ph *web.ProductHandler, Uh *web.UserHandler) *Router {
	r := mux.NewRouter()
	reg := prometheus.NewRegistry()
	reg.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
	)
	return &Router{Router: r, Metrics: reg, OrderHandler: Oh, ProductHandler: Ph, UserHandler: Uh}
}

func (r *Router) Init() {
	r.Router.HandleFunc("/api/user/{id}", r.UserHandler.Get)
	r.Router.HandleFunc("/api/users", r.UserHandler.List)

	r.Router.HandleFunc("/api/products", r.ProductHandler.List)

	r.Router.HandleFunc("/api/order/{id}", r.OrderHandler.Get)
	r.Router.HandleFunc("/api/orders", r.OrderHandler.List)

	r.Router.HandleFunc("/health", r.UserHandler.Health)
	r.Router.Handle("/metrics", promhttp.HandlerFor(r.Metrics, promhttp.HandlerOpts{}))

	http.Handle("/", r.Router)
	http.ListenAndServe(":8080", nil)
}
