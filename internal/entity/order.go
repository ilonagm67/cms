package entity

const (
	StateAwaitingOrderProducts State = "order_products"
	StateAwaitingOrderPayType  State = "order_paytype"
	StateAwaitingOrderDelivery State = "order_delivery"
	StateAwaitingOrderAddress  State = "order_address"
)

type Order struct {
	ID       int64
	UserID   int64
	Products []Product
	PayType  string
	Address  string
	Delivery string
}
