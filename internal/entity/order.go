package entity

const (
	StateAwaitingOrderProducts       State = "order_products"
	StateAwaitingOrderProductsWeight State = "order_products_weight"
	StateAwaitingOrderProductsCount  State = "order_products_count"
	StateAwaitingOrderPreDelivery    State = "order_pre_delivery"
	StateAwaitingOrderDelivery       State = "order_delivery"
	StateAwaitingOrderPayType        State = "order_paytype"
	StateAwaitingOrderAddress        State = "order_address"
)

type Order struct {
	ID       int64
	UserID   int64
	Products map[string]map[int]*Product
	PayType  string
	Address  string
	Delivery string
}
