package memory

import (
	"cms/internal/entity"
	"errors"
	"fmt"
)

type ProductRepository struct {
	products map[string]map[int]*entity.Product
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		products: make(map[string]map[int]*entity.Product, 0),
	}
}

func (repo *ProductRepository) Add(Product *entity.Product) error {
	if Product.Name == "" || Product.Weight == 0 {
		return fmt.Errorf("Product Name or Weight Not Found")
	}
	_, exists := repo.products[Product.Name]
	if !exists {
		repo.products[Product.Name] = make(map[int]*entity.Product)
	}
	repo.products[Product.Name][Product.Weight] = Product
	return nil
}

func (repo *ProductRepository) Delete(Name string, Weight int) error {
	if Name == "" || Weight == 0 {
		return fmt.Errorf("Name or Weight not Permitted!")
	}

	weights, ok := repo.products[Name]
	if !ok {
		return fmt.Errorf("ProductName: %s Not Found!", Name)
	}
	_, ok = weights[Weight]
	if !ok {
		return fmt.Errorf("ProductWeight: %d Not Found!", Weight)
	}

	delete(weights, Weight)

	if len(weights) == 0 {
		delete(repo.products, Name)
	}

	return nil
}

func (repo *ProductRepository) Get(Name string, Weight int) (*entity.Product, error) {
	i, ok := repo.products[Name][Weight]
	if ok {
		return i, nil
	}
	return nil, fmt.Errorf("Product Not Found, Name: %s, Weight: %d", Name, Weight)
}

func (repo *ProductRepository) List() (map[string]map[int]*entity.Product, error) {
	if len(repo.products) > 0 {
		return repo.products, nil
	}
	return nil, errors.New("No Items in Database Products")
}
