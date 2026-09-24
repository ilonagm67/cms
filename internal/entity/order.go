package entity

const (
	StateAwaitingOrderProducts             State = "order_products"
	StateAwaitingOrderProductsWeight       State = "order_products_weight"
	StateAwaitingOrderProductsCount        State = "order_products_count"
	StateAwaitingOrderProductsDeleteName   State = "order_products_delete_name"
	StateAwaitingOrderProductsDeleteWeight State = "order_products_delete_weight"
	StateAwaitingOrderPreDelivery          State = "order_pre_delivery"
	StateAwaitingOrderDelivery             State = "order_delivery"
	StateAwaitingOrderPayType              State = "order_paytype"
	StateAwaitingOrderAddress              State = "order_address"
)

type Order struct {
	UserID   int64
	Products map[string]map[int]*Product
	PayType  string
	Address  string
	Delivery string
}
