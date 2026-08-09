package entity

type Order struct {
	ID       uint64
	Customer User
	Products []Product
	PayType  string
	Address  string
	Delivery string
}
