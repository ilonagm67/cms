package entity

const (
	StateAwaitingProductName        State = "product_name"
	StateAwaitingProductDescription State = "product_description"
	StateAwaitingProductImage       State = "product_image"
	StateAwaitingProductWeight      State = "product_weight"
	StateAwaitingProductPrice       State = "product_price"
)

type Product struct {
	Name        string
	Description string
	Image       string
	Weight      int
	Count       int
	Price       float64
}
