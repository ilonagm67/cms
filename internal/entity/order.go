package entity

type Order struct {
	ID       int64
	Customer User
	Products []Product
	PayType  string
	Address  string
	Delivery string
}
