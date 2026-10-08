package entity

const (
	StateAwaitingProductList         State = "product_list"
	StateAwaitingProductListWeight   State = "product_list_weight"
	StateAwaitingProductName         State = "product_name"
	StateAwaitingProductDescription  State = "product_description"
	StateAwaitingProductImage        State = "product_image"
	StateAwaitingProductWeight       State = "product_weight"
	StateAwaitingProductPrice        State = "product_price"
	StateAwaitingProductDeleteName   State = "product_delete_name"
	StateAwaitingProductDeleteWeight State = "product_delete_weight"
)

type Product struct {
	Name        string
	Description string
	Image       string
	Weight      int
	Count       int
	Price       int
}
