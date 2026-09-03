package entity

type Order struct {
	ID       int64
	UserID   int64
	Products []Product
	PayType  string
	Address  string
	Delivery string
}
