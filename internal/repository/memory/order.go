package memory

import (
	"cms/internal/entity"
	"fmt"
)

type OrderRepository struct {
	orders map[int64]*entity.Order
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		orders: make(map[int64]*entity.Order, 0),
	}
}

func (repo *OrderRepository) Add(ID int64, Order *entity.Order) error {
	if Order.UserID == 0 {
		return fmt.Errorf("Order UserID Not Found!")
	}
	repo.orders[ID] = Order
	return nil
}

func (repo *OrderRepository) AddProduct(ID int64, Product *entity.Product) error {
	if Product.Name == "" || Product.Weight == 0 {
		return fmt.Errorf("Product Name or Weight Not Found")
	}
	if repo.orders[ID].Products == nil {
		productmap := make(map[string]map[int]*entity.Product)
		repo.orders[ID].Products = productmap
	}
	_, ok := repo.orders[ID].Products[Product.Name]
	if !ok {
		repo.orders[ID].Products[Product.Name] = make(map[int]*entity.Product)
	}
	repo.orders[ID].Products[Product.Name][Product.Weight] = Product
	return nil
}

func (repo *OrderRepository) DeleteProduct(ID int64, Name string, Weight int) error {
	order, ok := repo.orders[ID]
	if !ok {
		return fmt.Errorf("User %d Not Found!", ID)
	}
	if order.Products == nil {
		return fmt.Errorf("Map Not Found!")
	}
	weights, productExists := order.Products[Name]
	if !productExists {
		return fmt.Errorf("product category %q not found for customer %d", Name, ID)
	}

	if _, weightExists := weights[Weight]; !weightExists {
		return fmt.Errorf("weight %d not found under product %q", Weight, Name)
	}

	delete(weights, Weight)

	if len(weights) == 0 {
		delete(order.Products, Name)
	}
	return nil
}

func (repo *OrderRepository) Get(ID int64) (*entity.Order, error) {
	order, ok := repo.orders[ID]
	if !ok {
		return nil, fmt.Errorf("Order Not Found")
	}
	return order, nil
}

func (repo *OrderRepository) List() (map[int64]*entity.Order, error) {
	if len(repo.orders) > 0 {
		return repo.orders, nil
	}
	return nil, fmt.Errorf("No Items in Database Orders")
}
